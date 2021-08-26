package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
	"git.fractalqb.de/fractalqb/yacfg"
	"github.com/CmdrVasquess/watched"
	"github.com/CmdrVasquess/watched/jdir"

	"github.com/CmdrVasquess/edpc/internal"
	"github.com/CmdrVasquess/edpc/internal/cmds"
)

var (
	log    = qbsllm.New(qbsllm.Lnormal, "edpc", nil, nil)
	logCfg = c4hgol.Config(qbsllm.NewConfig(log),
		internal.LogCfg,
	)

	config = struct {
		Log         string
		JournalDir  string
		WatchLatest bool
		JDirWatch   jdir.Options
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
	flag.BoolVar(&config.WatchLatest, "j-latest", config.WatchLatest,
		`Don't wait for new journal file but start watching latest existing
journal`)
	flag.Int64Var(&config.JDirWatch.JSerial, "last-jeid", config.JDirWatch.JSerial,
		`Set last known journal event ID. 0 loads last JEID from DB.`)
	flag.StringVar(&config.Log, "log", "", c4hgol.LevelCfgDoc(nil))
	flag.Parse()
}

func main() {
	log.Infof("Running EDPC client v%d.%d.%d (%s #%d, %s)",
		internal.VMajor, internal.VMinor, internal.VPatch,
		internal.VQuality, internal.VBuildNo,
		runtime.Version(),
	)
	if err := cmds.Configure(&config); yacfg.IsCode(err, yacfg.ConfigQuery) {
		cmds.DumpConfig(os.Stdout, &config)
		os.Exit(0)
	} else if err != nil {
		log.Fatale(err)
	}
	flags()
	c4hgol.SetLevel(logCfg, config.Log, nil)
	edpc := internal.NewEDPC()
	watchED := jdir.NewEvents(config.JournalDir, edpc, &config.JDirWatch)
	var latestJournal string
	if config.WatchLatest {
		var err error
		latestJournal, err = jdir.NewestJournal(config.JournalDir)
		if err != nil {
			log.Fatale(err)
		}
	}
	go watchED.Start(latestJournal)
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Infos("shutting down…")
	// if err := bcp.SetLastJEID(watchED.LastJSerial()); err != nil {
	// 	log.Errore(err)
	// }
	watchED.Stop <- watched.Stop
	<-watchED.Stop
	log.Infoa("o7")
}