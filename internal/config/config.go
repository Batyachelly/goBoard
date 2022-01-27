package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	HTTP     HTTP
	Postgres Postgres
	General  General
}

func ParseConfig() (*Config, error) {
	cfg := new(Config)

	if err := env.Parse(&cfg.General); err != nil {
		return nil, fmt.Errorf("parse general config: %w", err)
	}

	if err := env.Parse(&cfg.HTTP); err != nil {
		return nil, fmt.Errorf("parse http config: %w", err)
	}

	if err := env.Parse(&cfg.Postgres); err != nil {
		return nil, fmt.Errorf("parse postgres config: %w", err)
	}

	return cfg, nil
}
