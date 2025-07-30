package http

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(ctx context.Context, config Config) (*chi.Mux, error) {
	router := chi.NewRouter()

	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	router.Use(middleware.Heartbeat("/ping"))

	return router, nil
}