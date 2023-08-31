package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/poggerr/avito/internal/app"
	"github.com/poggerr/avito/internal/config"
	"github.com/poggerr/avito/internal/logger"
	"github.com/poggerr/avito/internal/routers"
	"github.com/poggerr/avito/internal/server"
	"github.com/poggerr/avito/internal/storage"
	"log"
)

func main() {
	cfg := config.NewConf()

	db, err := sqlx.Connect("postgres", cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}
	sugaredLogger := logger.Initialize()

	strg := storage.NewStorage(db, cfg)

	newApp := app.NewApp(cfg, strg, sugaredLogger)

	r := routers.Router(newApp)
	server.Server(cfg.ServAddr, r)
}
