package internal

import (
	"errors"

	"github.com/CmdrVasquess/watched"
)

type EDPC struct {
}

func NewEDPC() *EDPC {
	return &EDPC{}
}

func (er *EDPC) OnJournalEvent(e watched.JounalEvent) error {
	// TODO
	return errors.New("NYI")
}

func (er *EDPC) OnStatusEvent(e watched.StatusEvent) error {
	// TODO
	return errors.New("NYI")
}

func (er *EDPC) Close() error {
	// TODO
	return errors.New("NYI")
}
