package http

import (
	"context"
	"log"
)

type App struct {
	Context  context.Context
	Config *Config
}

func New(ctx context.Context) *App {
	return &App{
		Context: ctx,
		Config: &Config{},
	}
}

func (a *App) SetConfig() {
	cfg, err := LoadConfig()

	if err != nil {
		log.Panic("Error loading config:", err)
	}

	a.Config = cfg
}