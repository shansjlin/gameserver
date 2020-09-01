package gameserver

import "sync"

type Game struct {
	routinesGroup sync.WaitGroup
	conf Config
	rpcCh <-chan RPC
	trans Transport
}

func NewGame(conf *Config, trans Transport) (*Game, error){

	g := &Game{
		conf:  *conf,
		rpcCh: trans.Consumer(),
		trans: trans,
	}
	return g, nil
}

func (g *Game) goFunc(f func()) {
	g.routinesGroup.Add(1)
	go func() {
		defer g.routinesGroup.Done()
		f()
	}()
}