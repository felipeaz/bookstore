package main

import (
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

	err = server.Start(dbService, logger)
	if err != nil {
		log.Fatal(err.Error())
	}
}
