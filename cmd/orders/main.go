package main

import (
	"bookstore/build/server/amqp/sender"
	"bookstore/infra/redis"
	"google.golang.org/grpc"
	"log"
	"os"

	"bookstore/build/server/orders/server"
	_log "bookstore/infra/logger"
	"bookstore/infra/mysql/orders/database"
	"bookstore/infra/mysql/service"
)

const (
	ServiceName = "orders Service"
)

func main() {
	db, err := database.Connect(
		os.Getenv("ORDERS_DB_USER"),
		os.Getenv("ORDERS_DB_PASSWORD"),
		os.Getenv("ORDERS_DB_HOST"),
		os.Getenv("ORDERS_DB_PORT"),
		os.Getenv("ORDERS_DB_DATABASE"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.CloseConnection(db)

	logger := _log.NewLogger(os.Getenv("LOG_FILE"), ServiceName)
	dbService := service.NewMySQLService(db, logger)

	cache, err := redis.NewCache(
		os.Getenv("ORDERS_REDIS_HOST"),
		os.Getenv("ORDERS_REDIS_PORT"),
		os.Getenv("ORDERS_REDIS_EXPIRE"),
		logger,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	grpcConn, err := grpc.Dial(os.Getenv("INVENTORY_GRPC_HOST"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer grpcConn.Close()

	queue, err := sender.CreateAMQP(os.Getenv("AMQP_SERVER_URL"), os.Getenv("AMQP_QUEUE_NAME"), logger)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer queue.QueueConn.Close()
	defer queue.QueueChannel.Close()

	err = server.Start(dbService, queue, grpcConn, cache, logger)
	if err != nil {
		log.Fatal(err.Error())
	}
}
