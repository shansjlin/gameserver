package gameserver

import (
	"io"
	"time"
)

type ProtocolVersion int
type ServerAddress string

type Config struct {
	HeartbeatTimeout time.Duration
	NotifyCh chan<- bool

	LogOutput io.Writer
	LogLevel string
}

func DefaultConfig() *Config {
	return &Config{
		HeartbeatTimeout:   1000 * time.Millisecond,
		LogLevel:           "DEBUG",
	}
}