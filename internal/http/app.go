package http

import (
	utils "api-template/internal/utils"
	"context"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

type App struct {
	Context  context.Context
	Env      *Env
	Injector *do.Injector
	Router   *chi.Mux
}

func New(ctx context.Context) *App {
	return &App{
		Context:  ctx,
		Env:      &Env{},
		Injector: nil,
		Router:   nil,
	}
}

func (a *App) SetConfig() {
	cfg, err := LoadConfig()

	if err != nil {
		log.Panic("Error loading config:", err)
	}

	a.Env = cfg
}

func (a *App) SetInjector() {
	injector, err := CreateInjector(a.Context, *a.Env)

	if err != nil {
		log.Panic("Error creating injector:", err)
	}

	a.Injector = injector
}

func (a *App) SetRouter() {
	router, err := CreateRouter(a.Context, *a.Env)

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

	dbUrl := a.Env.ConnectionString()

	if err := utils.RunMigrations(dbUrl); err != nil {
		log.Panic("Error running migrations:", err)
	}

	log.Println("Application started successfully")
}
