package main

import (
	"bookstore/infra/redis"
	"log"
	"os"

	"bookstore/build/server/inventory/server"
	_log "bookstore/infra/logger"
	"bookstore/infra/mysql/inventory/database"
	"bookstore/infra/mysql/service"
)

const (
	ServiceName = "inventory Service"
)

func main() {
	db, err := database.Connect(
		os.Getenv("INVENTORY_DB_USER"),
		os.Getenv("INVENTORY_DB_PASSWORD"),
		os.Getenv("INVENTORY_DB_HOST"),
		os.Getenv("INVENTORY_DB_PORT"),
		os.Getenv("INVENTORY_DB_DATABASE"),
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

	err = server.Start(dbService, cache, logger)
	if err != nil {
		log.Fatal(err.Error())
	}
}
