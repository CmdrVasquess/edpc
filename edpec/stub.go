package edpec

import (
	"net/rpc"

	"github.com/CmdrVasquess/edpc/edpec/api"
)

type Config struct {
	Addr string
}

type Stub struct {
	cfg      *Config
	client   *rpc.Client
	APIToken string
}

func NewStub(cfg *Config) (*Stub, error) {
	client, err := rpc.DialHTTP("tcp", cfg.Addr)
	if err != nil {
		return nil, err
	}
	return &Stub{
		cfg:    cfg,
		client: client,
	}, nil
}

func (edpec *Stub) Close() error {
	return edpec.client.Close()
}

func (edpec *Stub) Dock(ssys string, addr int64, port string) (*api.DockReply, error) {
	args := &api.DockRequest{
		System: ssys,
		Addr:   addr,
		Port:   port,
	}
	reply := new(api.DockReply)
	err := edpec.client.Call(svc+"Dock", args, reply)
	return reply, err
}

const svc = "EdpEc."
