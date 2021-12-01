package server

import (
	"bookstore/build/server/amqp/sender"
	"bookstore/build/server/orders/router"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/orders/order/handler"
	"bookstore/internal/app/domain/orders/order/module"
	"bookstore/internal/app/domain/orders/order/repository"
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
	cRepository := repository.NewOrderRepository(dbService)
	cModule := module.NewOrderModule(cRepository, queue, grpcConn, cache, log)
	cHandler := handler.NewOrderHandler(cModule)

	return router.Build(cHandler)
}
