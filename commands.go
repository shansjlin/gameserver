package gameserver

type RPCHeader struct {
	ProtocolVersion ProtocolVersion
}

type WithRPCHeader interface {
	GetRPCHeader() RPCHeader
}

type RPCRequest struct {
	RPCHeader
	Payload []byte
}

func (r *RPCRequest) GetRPCHeader() RPCHeader {
	return r.RPCHeader
}

type RPCRequestRsp struct {
	RPCHeader
	Payload []byte
}

func (r *RPCRequestRsp) GetRPCHeader() RPCHeader {
	return r.RPCHeader
}
