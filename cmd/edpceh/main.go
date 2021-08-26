package main

import (
	"flag"
	"os"
	"runtime"

	"git.fractalqb.de/fractalqb/yacfg"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
	"github.com/CmdrVasquess/edpc/internal"
	"github.com/CmdrVasquess/edpc/internal/cmds"
	"github.com/CmdrVasquess/watched/edeh/edehnet"
)

var (
	log    = qbsllm.New(qbsllm.Lnormal, "edpceh", nil, nil)
	logCfg = c4hgol.Config(qbsllm.NewConfig(log),
		internal.LogCfg,
	)

	config = struct {
		Log     string
		EDEHNet edehnet.Receiver
	}{}
)

func flags() {
	flag.StringVar(&config.Log, "log", "", c4hgol.LevelCfgDoc(nil))
	flag.Parse()
}

func main() {
	log.Infof("Running EDPC-edeh v%d.%d.%d (%s #%d, %s)",
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
	for {
		if err := config.EDEHNet.Run(edpc); err != nil {
			log.Errore(err)
		}
	}
}
