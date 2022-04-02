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
	var evt journal.Docked
	err := json.Unmarshal(rawe.Event, &evt)
	if err != nil {
		return err
	}
	log.Warna("TODO: docked `at` `in`", evt.StationName, evt.StarSystem)
	return err
}
