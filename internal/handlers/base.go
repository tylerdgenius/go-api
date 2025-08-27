package handlers

import (
	"api-template/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

type BaseHandler struct {
	sampleService services.ISampleService
}

func RegisterBaseHandlerGroup(r chi.Router, prefix string, i *do.Injector) {
	handlers := RegisterHandlers(i)

	r.Route(prefix, func(r chi.Router) {
		for _, h := range handlers {
			r.Method(h.method, h.pattern, h.handler)
		}
	})
}
