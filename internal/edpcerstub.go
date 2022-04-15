package internal

import (
	"context"
	"errors"
	"time"

	"github.com/CmdrVasquess/edpc/edpcrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EDPCer struct {
	conn    *grpc.ClientConn
	clt     edpcrpc.EDPCerClient
	timeout time.Duration
	apiKey  string
}

func NewEDPCer(addr string) (*EDPCer, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &EDPCer{
		conn:    conn,
		clt:     edpcrpc.NewEDPCerClient(conn),
		timeout: 5 * time.Second,
	}, nil
}

func (er *EDPCer) Close() {
	er.conn.Close()
	// TODO: Close EDPCer
}

func (er *EDPCer) Commander(fid, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), er.timeout)
	defer cancel()
	_, err := er.clt.Commander(ctx, &edpcrpc.CommanderRequest{
		Hdr:  er.rqHeader(fid),
		Cmdr: name,
	})
	if err != nil {
		return err
	}
	// TODO: Handle response
	return errors.New("NYI: EDPCer Commander: return")
}

func (er *EDPCer) Docked(fid string, addr uint64, ssys, port string) error {
	ctx, cancel := context.WithTimeout(context.Background(), er.timeout)
	defer cancel()
	_, err := er.clt.Docked(ctx, &edpcrpc.DockedRequest{
		Hdr:    er.rqHeader(fid),
		Addr:   addr,
		System: ssys,
		Port:   port,
	})
	if err != nil {
		return err
	}
	// TODO: Handle response
	return errors.New("NYI: EDPCer.Docked")
}

func (er *EDPCer) rqHeader(fid string) *edpcrpc.RequestHeader {
	return &edpcrpc.RequestHeader{
		AccessToken: er.apiKey,
		Fid:         fid,
	}
}
