package api

type DockRequest struct {
	RequestBase
	System string
	Addr   int64
	Port   string
}

type DockReply struct {
	Discos []int32
}
