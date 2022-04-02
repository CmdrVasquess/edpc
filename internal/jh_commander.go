package internal

import (
	"encoding/json"

	"github.com/CmdrVasquess/stated/journal"
	"github.com/CmdrVasquess/watched"
)

func init() {
	journalHandlers[journal.CommanderEvent.String()] = jehCommander
}

func jehCommander(edpc *EDPC, rawe watched.JounalEvent) error {
	var evt journal.Commander
	err := json.Unmarshal(rawe.Event, &evt)
	if err != nil {
		return err
	}
	err = edpc.Cmdr.SwitchTo(edpc.App, evt.FID, evt.Name)
	return err
}
