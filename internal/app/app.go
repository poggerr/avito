package app

import (
	"github.com/poggerr/avito/internal/config"
	"github.com/poggerr/avito/internal/storage"
	"go.uber.org/zap"
)

type App struct {
	cfg           *config.Config
	strg          *storage.Storage
	sugaredLogger *zap.SugaredLogger
}

func NewApp(cfg *config.Config, strg *storage.Storage, log *zap.SugaredLogger) *App {
	return &App{
		cfg:           cfg,
		strg:          strg,
		sugaredLogger: log,
	}
}
