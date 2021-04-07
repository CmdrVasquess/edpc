package main

import (
	"github.com/CmdrVasquess/edpc/edpec"
	"github.com/CmdrVasquess/stated"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/CmdrVasquess/watched"
	"github.com/CmdrVasquess/watched/jdir"
)

type Config struct {
	Stated   stated.Config
	DataDir  string
	JDir     string
	JDirOpts jdir.Options
	EdpEc    edpec.Config
}

var (
	edstate *stated.EDState
	watcher *jdir.Events
	changes = make(chan stated.ChangeEvent, 32)
	edpeclt *edpec.Stub
)

func runEventClient(cfg *Config) {
	log.Infoa("using `data dir`", cfg.DataDir)
	cfg.Stated.CmdrFile = stated.CmdrFile{Dir: cfg.DataDir}.Filename
	edstate = stated.NewEDState(&cfg.Stated)
	edstate.Notify = []chan<- stated.ChangeEvent{changes}
	watcher = jdir.NewEvents(cfg.JDir, edstate, &cfg.JDirOpts)
	var err error
	edpeclt, err = edpec.NewStub(&cfg.EdpEc)
	if err != nil {
		log.Fatale(err)
	}
	go watcher.Start("")
	go changeHandler(changes)
}

func stopEventClient() {
	watcher.Stop <- watched.Stop
	<-watcher.Stop
	edstate.Close()
	changes <- stated.ChangeEvent{}
	<-changes
	edpeclt.Close()
}

func changeHandler(chgs <-chan stated.ChangeEvent) {
	log.Infos("running change hanlder")
	for chg := range chgs {
		if chg.Change == 0 && chg.Event == nil {
			break
		}
		log.Tracea("`change`", chg)
		switch chg.Event.Event() {
		case journal.DockedEvent.String():
			edpeclt.Dock(
				edstate.Loc.System().Name,
				int64(edstate.Loc.System().Addr),
				edstate.Loc.Port().Name,
			)
		}
	}
	log.Infos("exit change handler")
	close(changes)
}
