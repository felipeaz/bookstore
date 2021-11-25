package server

import (
	"bookstore/build/server/inventory/router"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/inventory/books/handler"
	"bookstore/internal/app/domain/inventory/books/module"
	"bookstore/internal/app/domain/inventory/books/repository"
	"bookstore/internal/app/logger"
)

// Start initialize the webservice,
func Start(dbService database.GORMServiceInterface, log logger.LogInterface) (err error) {
	bRepository := repository.NewBookRepository(dbService)
	bModule := module.NewBookModule(bRepository, log)
	bHandler := handler.NewBookHandler(bModule)

	return router.Build(bHandler)
}
