package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

type BaseHandler struct {
}

func NewBaseHandler(r chi.Router, i *do.Injector) *BaseHandler {
	handler := &BaseHandler{}

	return handler
}
