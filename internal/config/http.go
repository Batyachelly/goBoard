package config

import "time"

type HTTP struct {
	Addr         string        `env:"HTTP_ADDR"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"15s"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT"  envDefault:"15s"`
}
