package server

import (
	"bookstore/build/server/account/router"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/account/auth"
	"bookstore/internal/app/domain/account/users/handler"
	"bookstore/internal/app/domain/account/users/module"
	"bookstore/internal/app/domain/account/users/repository"
	"bookstore/internal/app/logger"
)

// Start initialize the webservice,
func Start(
	dbService database.GORMServiceInterface,
	cache database.CacheInterface,
	apiGatewayAuth auth.Interface,
	log logger.LogInterface,
) (err error) {
	accountRepository := repository.NewAccountRepository(dbService)
	accountModule := module.NewAccountModule(accountRepository, apiGatewayAuth, cache, log)
	accountHandler := handler.NewAccountHandler(accountModule)
	return router.Build(accountHandler)
}
