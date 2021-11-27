package main

import (
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
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_EXPIRE"),
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
	defer func(grpcConnection *grpc.ClientConn) {
		err := grpcConnection.Close()
		if err != nil {

		}
	}(grpcConn)

	err = server.Start(dbService, grpcConn, cache, logger)
	if err != nil {
		log.Fatal(err.Error())
	}
}
