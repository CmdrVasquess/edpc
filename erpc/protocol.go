package erpc

const (
	ProtoName    = "EDPCer"
	ProtoVersion = 0
)

type Requester interface {
	RqHeader() *RequestHeader
}

type RequestHeader struct {
	ProtoVer    int    `cbor:"1,keyasint"`
	AccessToken string `cbor:"2,keyasint"`
	FID         string `cbor:"3,keyasint"`
}

func (h *RequestHeader) RqHeader() *RequestHeader { return h }

type Responder interface {
	RspHeader() *ResponseHeader
}

type ResponseHeader struct {
	ProtoVer int    `cbor:"1,keyasint"`
	Redirect string `cbor:"2,keyasint,omitempty"`
}

func (h *ResponseHeader) RspHeader() *ResponseHeader { return h }
