package handlers

import (
	"api-template/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

type BaseHandler struct {
	sampleService services.ISampleService
}

func NewBaseHandler(r chi.Router, i *do.Injector) *BaseHandler {
	handler := &BaseHandler{
		sampleService: do.MustInvoke[services.ISampleService](i),
	}

	r.Route("/sample", func(r chi.Router) {
		r.Get("/", handler.SampleHandler)
	})

	return handler
}
