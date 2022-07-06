package erpc

type Eventer interface {
	Requester
	EvtHeader() *EDEvent
}

type EDEvent struct {
	RequestHeader
	TEvt int64 `cbor:"4,keyasint"`
}

func (e *EDEvent) EvtHeader() *EDEvent { return e }

type CommanderEvent struct {
	EDEvent
	Cmdr string
}

type DockedEvent struct {
	EDEvent
	Addr   uint64 `cbor:",omitempty"`
	System string
	Port   string
}
