package main

import (
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
		os.Getenv("inventory_DB_USER"),
		os.Getenv("inventory_DB_PASSWORD"),
		os.Getenv("inventory_DB_HOST"),
		os.Getenv("inventory_DB_PORT"),
		os.Getenv("inventory_DB_DATABASE"),
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
