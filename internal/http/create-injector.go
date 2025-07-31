package http

import (
	"context"
	"log"

	"github.com/samber/do"
)

func CreateInjector(ctx context.Context, config Config) (*do.Injector, error) {
	injector := do.New()

	log.Println("Injector initialization successful")

	return injector, nil
}
