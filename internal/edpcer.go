package internal

import (
	"crypto/tls"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/lucas-clemente/quic-go"

	"github.com/CmdrVasquess/edpc/erpc"
)

type EDPCer struct {
	conn    quic.Connection
	enc     cbor.EncMode
	dec     cbor.DecMode
	accsTok string
}

func NewEDPCer(acstk string) (e *EDPCer, err error) {
	tags, err := erpc.CBORTags()
	if err != nil {
		return nil, err
	}
	// TODO Need more inits
	e = &EDPCer{accsTok: acstk}
	e.enc, err = cbor.EncOptions{}.EncModeWithTags(tags)
	if err != nil {
		return nil, err
	}
	e.dec, err = cbor.DecOptions{}.DecModeWithTags(tags)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (er *EDPCer) Open(addr string, insec bool) (err error) {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	er.conn, err = quic.DialAddr(addr, tlsConf, nil)
	return err
}

func (er *EDPCer) Close() error { return nil }

func (er *EDPCer) Commander(t time.Time, fid, name string) error {
	_, err := er.rpcEvent(fid, t, &erpc.CommanderEvent{
		Cmdr: name,
	})
	return err
}

func (er *EDPCer) Docked(t time.Time, fid, ssys, port string) error {
	_, err := er.rpcEvent(fid, t, &erpc.DockedEvent{
		System: ssys,
		Port:   port,
	})
	return err
}

func (er *EDPCer) stream() (quic.Stream, error) {
	return er.conn.OpenStream()
}

func (er *EDPCer) rpcEvent(fid string, t time.Time, e erpc.Eventer) (any, error) {
	e.EvtHeader().TEvt = t.Unix()
	return er.rpc(fid, e)
}

func (er *EDPCer) rpc(fid string, data erpc.Requester) (any, error) {
	hdr := data.RqHeader()
	hdr.ProtoVer = erpc.ProtoVersion
	hdr.AccessToken = er.accsTok
	hdr.FID = fid
	s, err := er.stream()
	if err != nil {
		return nil, err
	}
	// TODO Set deadline
	err = er.enc.NewEncoder(s).Encode(data)
	if err != nil {
		return nil, err
	}
	var resp any
	// TODO Set deadline
	err = er.dec.NewDecoder(s).Decode(&resp)
	s.Close()
	if err != nil {
		return nil, err
	}
	if r, ok := resp.(erpc.Responder); ok {
		err = er.redirect(r)
	}
	return resp, err
}

func (er *EDPCer) redirect(h erpc.Responder) error {
	raddr := h.RspHeader().Redirect
	if raddr == "" {
		return nil
	}
	// https://github.com/lucas-clemente/quic-go/issues/3270
	err := er.conn.CloseWithError(0, "redirected")
	if err != nil {
		return err
	}
	return er.Open(raddr, false)
}
