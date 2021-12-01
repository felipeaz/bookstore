package server

import (
	"bookstore/build/server/amqp/sender"
	"bookstore/build/server/orders/router"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/orders/items/handler"
	"bookstore/internal/app/domain/orders/items/module"
	"bookstore/internal/app/domain/orders/items/repository"
	"bookstore/internal/app/logger"
	"google.golang.org/grpc"
)

// Start initialize the webservice,
func Start(
	dbService database.GORMServiceInterface,
	queue *sender.RabbitMQ,
	grpcConn *grpc.ClientConn,
	cache database.CacheInterface,
	log logger.LogInterface) (err error) {
	cRepository := repository.NewItemRepository(dbService)
	cModule := module.NewItemModule(cRepository, queue, grpcConn, cache, log)
	cHandler := handler.NewItemHandler(cModule)

	return router.Build(cHandler)
}
