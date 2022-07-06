package main

import (
	"flag"
	"os"
	"runtime"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qblog"
	"github.com/CmdrVasquess/edpc/internal"
	"github.com/CmdrVasquess/edpc/internal/cmds"
	"github.com/CmdrVasquess/watched/edeh/edehnet"
)

var (
	log    = qblog.New("edpceh")
	logCfg = c4hgol.NewLogGroup(log, "", internal.LogCfg)

	config = struct {
		Log         string
		EDEHNet     edehnet.Receiver
		ERAddress   string
		AccessToken string
		InsecureER  bool
	}{
		EDEHNet: edehnet.Receiver{
			Listen: ":1337",
		},
	}
)

func flags() {
	flag.StringVar(&config.ERAddress, "r", config.ERAddress,
		`Set address of EDPCer server`,
	)
	flag.StringVar(&config.EDEHNet.Listen, "l", config.EDEHNet.Listen,
		`Set listening address`,
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
	log.Infof("Running EDPC-edeh v%d.%d.%d (%s #%d, %s)",
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
	for {
		if err := config.EDEHNet.Run(edpc); err != nil {
			log.Errore(err)
		}
	}
	defer edpc.Close()
}
