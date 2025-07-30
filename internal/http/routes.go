package http

import (
	"api-template/internal/handlers"
	"context"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(ctx context.Context, config Config) (*chi.Mux, error) {
	router := chi.NewRouter()

	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	router.Use(middleware.Heartbeat("/ping"))

	log.Println("Setting up routes...")

	router.Get("/test", handlers.TestHandler)

	return router, nil
}