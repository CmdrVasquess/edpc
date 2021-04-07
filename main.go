package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
	"github.com/CmdrVasquess/stated"
	"github.com/CmdrVasquess/watched/jdir"
)

const apptag = "edpcp"

var (
	config = Config{
		Stated: stated.Config{
			ShutdownLogsOut: true,
		},
		JDirOpts: jdir.Options{
			SerialIndependent: []string{
				"Fileheader",
				"Commander",
				"Shutdown",
			},
		},
	}

	log    = qbsllm.New(qbsllm.Lnormal, apptag, nil, nil)
	logCfg = c4hgol.Config(qbsllm.NewConfig(log),
		stated.LogCfg,
	)
)

var (
	fLog      string
	fCfgList  string
	fCfgPrint bool
)

func main() {
	flag.StringVar(&fCfgList, "cfg", "", "list of config files to read")
	flag.BoolVar(&fCfgPrint, "cfg-print", false, "print config")
	flag.StringVar(&fLog, "log", "", c4hgol.LevelCfgDoc(nil))
	flag.Parse()
	if fLog == "?" {
		fmt.Println(apptag + " loggers:")
		c4hgol.ListLogs(logCfg, os.Stdout, "  - ")
		return
	}
	c4hgol.SetLevel(logCfg, fLog, nil)
	config.DataDir = findDataDir()
	config.JDir = findJournalDir()
	readConfigs()
	runEventClient(&config)
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Infof("shutting down %s…", apptag)
	stopEventClient()
	log.Infos("bye!")
}

func findJournalDir() string {
	res, err := jdir.FindJournalDir()
	if err != nil {
		log.Fatale(err)
	}
	return res
}

func readConfigs() {
	cfgs := filepath.SplitList(fCfgList)
	for _, cfg := range cfgs {
		rd, err := os.Open(cfg)
		if err != nil {
			log.Fatale(err)
		}
		if err = json.NewDecoder(rd).Decode(&config); err != nil {
			log.Fatale(err)
		}
		rd.Close()
	}
	if fCfgPrint {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(&config)
		os.Exit(0)
	}
}
