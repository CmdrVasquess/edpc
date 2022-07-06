package internal

import (
	"encoding/json"

	"github.com/CmdrVasquess/stated/journal"
	"github.com/CmdrVasquess/watched"
)

func init() {
	journalHandlers[journal.DockedEvent.String()] = jehDocked
}

func jehDocked(edpc *EDPC, rawe watched.JounalEvent) error {
	if err := edpc.needCmdr(); err != nil {
		return err
	}
	var evt journal.Docked
	err := json.Unmarshal(rawe.Event, &evt)
	if err != nil {
		return err
	}
	log.Warnv("TODO: docked `at` `in`", evt.StationName, evt.StarSystem)
	err = edpc.er.Docked(evt.Time, edpc.Cmdr.FID,
		//evt.SystemAddress,
		evt.StarSystem,
		evt.StationName,
	)
	return err
}
