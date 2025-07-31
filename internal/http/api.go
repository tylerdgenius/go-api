package http

import (
	"context"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

type App struct {
	Context  context.Context
	Config *Config
	Injector *do.Injector
	Router *chi.Mux
}

func New(ctx context.Context) *App {
	return &App{
		Context: ctx,
		Config: &Config{},
		Injector: nil,
	}
}

func (a *App) SetConfig() {
	cfg, err := LoadConfig()

	if err != nil {
		log.Panic("Error loading config:", err)
	}

	a.Config = cfg
}

func (a *App) SetInjector() {
	injector, err := NewInjector(a.Context, *a.Config)

	if err != nil {
		log.Panic("Error creating injector:", err)
	}

	a.Injector = injector
}

func (a *App) SetRouter() {
	router, err := NewRouter(a.Context, *a.Config)

	if err != nil {
		log.Panic("Error creating router:", err)
	}

	a.Router = router
}

func (a *App) Shutdown() {
	log.Println("Shutting down application...")

	if a.Injector != nil {
		if err := a.Injector.Shutdown(); err != nil {
			log.Panic("Error closing injector:", err)
		}
	}
}

func (a *App) Run() {
	a.SetConfig()
	a.SetInjector()
	a.SetRouter()

	log.Println("Application started successfully")
}