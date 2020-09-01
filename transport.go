package gameserver

import "io"

type RPCResponse struct {
	Response interface{}
	Error    error
}

type RPC struct {
	Command  interface{}
	Reader   io.Reader
	RespChan chan<- RPCResponse
}

func (r *RPC) Respond(resp interface{}, err error) {
	r.RespChan <- RPCResponse{resp, err}
}

type Transport interface {
	Consumer() <-chan RPC
	SetHeartbeatHandler(cb func(rpc RPC))
}