package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qblog"
	"github.com/CmdrVasquess/watched/jdir"

	"github.com/CmdrVasquess/edpc/internal"
	"github.com/CmdrVasquess/edpc/internal/cmds"
)

var (
	log    = qblog.New("edpc")
	logCfg = c4hgol.NewLogGroup(log, "", internal.LogCfg)

	config = struct {
		Log         string
		JournalDir  string
		JDirWatch   jdir.Options
		ERAddress   string
		AccessToken string
		InsecureER  bool
	}{
		JournalDir: findJournals(),
	}
)

func findJournals() string {
	dir, err := jdir.FindJournalDir()
	if err != nil {
		return "."
	}
	return dir
}

func flags() {
	flag.StringVar(&config.JournalDir, "j", config.JournalDir,
		`Explicitly set ED directory with journal files`)
	flag.Int64Var(&config.JDirWatch.JSerial, "last-jeid", config.JDirWatch.JSerial,
		`Set last known journal event ID. 0 loads last JEID from DB.`)
	flag.StringVar(&config.ERAddress, "r", config.ERAddress,
		`Set address of EDPCer server`,
	)
	flag.StringVar(&config.Log, "log", "", c4hgol.FlagDoc())
	cfgDump := flag.Bool("cfg-dump", false, "Dump current configuration to stdout")
	flag.Parse()
	if *cfgDump {
		cmds.DumpConfig(os.Stdout, &config)
		os.Exit(0)
	}
}

func main() {
	log.Infof("Running EDPC client v%d.%d.%d (%s #%d, %s)",
		internal.VMajor, internal.VMinor, internal.VPatch,
		internal.VQuality, internal.VBuildNo,
		runtime.Version(),
	)
	if err := cmds.Configure(&config); err != nil {
		log.Fatale(err)
	}
	flags()
	c4hgol.Configure(logCfg, config.Log, true)
	edpc, err := internal.NewEDPC(config.ERAddress, config.AccessToken, config.InsecureER)
	if err != nil {
		log.Fatale(err)
	}
	log.Infov("Local data in `dir`", edpc.App.LocalData())
	log.Infov("Roaming data in `dir`", edpc.App.RoamingData())
	log.Infov("EDPC `event receiver`", config.ERAddress)
	watchED := jdir.NewEvents(config.JournalDir, edpc, &config.JDirWatch)
	var latestJournal string
	go watchED.Start(latestJournal)
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Infos("shutting down…")
	watchED.Stop()
	edpc.Close()
	log.Infos("o7")
}
