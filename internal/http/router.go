package http

import (
	"api-template/internal/handlers"
	"context"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samber/do"
)

func CreateRouter(ctx context.Context, env *Env, injector *do.Injector) (*chi.Mux, error) {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.Heartbeat("/ping"))

	log.Println("Setting up routes...")

	// Example: Register handler group with custom prefix
	handlers.RegisterBaseHandlerGroup(router, "/api", injector)

	log.Println("Routes initialized successfully")

	return router, nil
}
