package handlers

import (
	"api-template/internal/services"
	"net/http"

	"github.com/samber/do"
)

type RegisteredHandler struct {
	pattern string
	method string
	handler http.HandlerFunc
}

func RegisterHandlers(i *do.Injector) []*RegisteredHandler {
	handler := &BaseHandler{
		sampleService: do.MustInvoke[services.ISampleService](i),
	}

	return []*RegisteredHandler{
		{
			pattern: "/sample",
			method:  http.MethodGet,
			handler: handler.SampleHandler,
		},
	}
}