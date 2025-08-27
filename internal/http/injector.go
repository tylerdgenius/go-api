package http

import (
	"api-template/internal/repository"
	"api-template/internal/services"
	"context"
	"log"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func CreateInjector(ctx context.Context, env Env) (*do.Injector, error) {
	injector := do.New()

	log.Println("Injector initialization successful")

	do.Provide(injector, func(i *do.Injector) (*gorm.DB, error) {
		db, err := NewDatabase(env)

		if err != nil {
			log.Panic("Error creating database:", err)
		}

		return db.Connect(ctx)
	})

	do.Provide(injector, func(i *do.Injector) (repository.ISampleRepository, error) {
		return repository.NewSampleRepository(do.MustInvoke[*gorm.DB](i)), nil
	})

	do.Provide(injector, func(i *do.Injector) (services.ISampleService, error) {
		return services.NewSampleService(do.MustInvoke[repository.ISampleRepository](i)), nil
	})

	return injector, nil
}
