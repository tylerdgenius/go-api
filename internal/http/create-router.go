package http

import (
	"api-template/internal/handlers"
	"context"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateRouter(ctx context.Context, config Config) (*chi.Mux, error) {
	router := chi.NewRouter()

	// Add basic middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.Heartbeat("/ping"))

	log.Println("Setting up routes...")

	// Set up your routes
	router.Get("/test", handlers.TestHandler)

	log.Println("Routes initialized successfully")

	return router, nil
}
