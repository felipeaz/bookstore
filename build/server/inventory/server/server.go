package server

import (
	"bookstore/build/server/inventory/grpc"
	"bookstore/build/server/inventory/router"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/inventory/books/handler"
	"bookstore/internal/app/domain/inventory/books/module"
	"bookstore/internal/app/domain/inventory/books/repository"
	"bookstore/internal/app/logger"
	"log"
)

// Start initialize the webservice,
func Start(
	dbService database.GORMServiceInterface,
	cache database.CacheInterface,
	lg logger.LogInterface) (err error) {
	bRepository := repository.NewBookRepository(dbService)
	bModule := module.NewBookModule(bRepository, cache, lg)
	bHandler := handler.NewBookHandler(bModule)

	go func() {
		err = grpc.Start(bModule)
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	return router.Build(bHandler)
}
