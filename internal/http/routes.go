package http

import (
	"context"

	"github.com/go-chi/chi/v5"
)

func NewRouter(ctx context.Context, config Config) (*chi.Mux, error) {
	router := chi.NewRouter()

	return router, nil
}