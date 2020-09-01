package main

import (
	"fmt"
	"net"
	"os"
	"time"
)
import "github.com/shansjlin/gameserver"

func main() {
	cfg := gameserver.DefaultConfig()
	fmt.Println(cfg)
	bindAddr := "127.0.0.1:8080"
	addr, _ := net.ResolveIPAddr("tcp", bindAddr)
	transport, _ := gameserver.NewTCPTransport(bindAddr, addr, 1, 10 * time.Second, os.Stderr)
	gameserver.NewGame(cfg, transport)
	time.Sleep(3600 * time.Second)
}
