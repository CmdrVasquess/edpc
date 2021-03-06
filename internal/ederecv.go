package internal

import (
	"errors"
	"os/user"

	"git.fractalqb.de/fractalqb/pack/ospath"
	"github.com/CmdrVasquess/watched"
)

type EDPC struct {
	App  ospath.App
	Cmdr Commander
	er   *EDPCer
	// TODO Last Event Serial Number
}

func NewEDPC(addr, accsTok string, insec bool) (*EDPC, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	edpc := &EDPC{
		App: ospath.NewApp(
			ospath.NewDesktopUserApp(user.Username),
			"fqb", "edpc",
		),
	}
	edpc.er, err = NewEDPCer(addr)
	if err != nil {
		return nil, err
	}
	edpc.er.Open(addr, insec)
	return edpc, err
}

type journalHandler = func(*EDPC, watched.JounalEvent) error

var journalHandlers = make(map[string]journalHandler)

func (edpc *EDPC) OnJournalEvent(e watched.JounalEvent) error {
	evt, err := watched.PeekEvent(e.Event)
	if err != nil {
		return err
	}
	h := journalHandlers[evt]
	if h != nil {
		err = h(edpc, e)
		if err == nil {
			edpc.Cmdr.System.ESeq = e.Serial
		}
	}
	return err
}

func (edpc *EDPC) OnStatusEvent(e watched.StatusEvent) error {
	// TODO
	return errors.New("NYI")
}

func (edpc *EDPC) Close() error {
	err := edpc.Cmdr.SwitchTo(edpc.App, "", "")
	if err != nil {
		log.Errore(err)
	}
	return err
}

func (edpc *EDPC) needCmdr() error {
	if edpc.Cmdr.Unknown() {
		return errors.New("No active commander")
	}
	return nil
}
