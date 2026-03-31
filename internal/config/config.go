package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v8"
	"strings"
)

type Config struct {
	ServAddr string `env:"RUN_ADDRESS"`
	DB       string `env:"DATABASE_URI"`
}

func NewConf() *Config {
	var cfg Config

	flag.StringVar(&cfg.ServAddr, "a", ":8080", "write down server")
	flag.StringVar(
		&cfg.DB,
		"d",
		"host=db user=avito password=password dbname=avito sslmode=disable",
		"write down db")
	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	cfg.ServAddr = strings.TrimSpace(cfg.ServAddr)
	cfg.DB = strings.TrimSpace(cfg.DB)
	if cfg.ServAddr == "" {
		cfg.ServAddr = ":8080"
	}
	if cfg.DB == "" {
		cfg.DB = "host=db user=avito password=password dbname=avito sslmode=disable"
	}
	return &cfg
}
