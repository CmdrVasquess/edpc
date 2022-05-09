package edpcrpc

import (
	"context"
	"crypto/x509"
	_ "embed"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:embed edpc.cert
var rootCertPEM []byte

type EDPCer struct {
	conn    *grpc.ClientConn
	clt     EDPCerClient
	timeout time.Duration
	accsTok string
}

func NewEDPCer(addr, accsTok string, insec bool) (*EDPCer, error) {
	var conn *grpc.ClientConn
	var err error
	if insec {
		conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		certs := x509.NewCertPool()
		certs.AppendCertsFromPEM(rootCertPEM)
		conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(
			credentials.NewClientTLSFromCert(certs, ""),
		))
	}
	if err != nil {
		return nil, err
	}
	return &EDPCer{
		conn:    conn,
		clt:     NewEDPCerClient(conn),
		timeout: 5 * time.Second,
		accsTok: accsTok,
	}, nil
}

func (er *EDPCer) Close() {
	er.conn.Close()
	// TODO: Close EDPCer
}

func (er *EDPCer) Commander(t time.Time, fid, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), er.timeout)
	defer cancel()
	_, err := er.clt.Commander(ctx, &CommanderRequest{
		Hdr:  er.rqHeader(t, fid),
		Cmdr: name,
	})
	if err != nil {
		return err
	}
	// TODO: Handle response
	return nil
}

func (er *EDPCer) Docked(t time.Time, fid string, addr uint64, ssys, port string) error {
	ctx, cancel := context.WithTimeout(context.Background(), er.timeout)
	defer cancel()
	_, err := er.clt.Docked(ctx, &DockedRequest{
		Hdr:    er.rqHeader(t, fid),
		Addr:   addr,
		System: ssys,
		Port:   port,
	})
	if err != nil {
		return err
	}
	// TODO: Handle response
	return nil
}

func (er *EDPCer) rqHeader(t time.Time, fid string) *RequestHeader {
	return &RequestHeader{
		AccessToken: er.accsTok,
		Fid:         fid,
		Tevt:        timestamppb.New(t),
	}
}
