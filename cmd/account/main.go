package main

import (
	"log"
	"os"

	"bookstore/build/server/account/server"
	"bookstore/infra/auth"
	"bookstore/infra/auth/http/client"
	"bookstore/infra/auth/http/request"
	_log "bookstore/infra/logger"
	"bookstore/infra/mysql/account/database"
	"bookstore/infra/mysql/service"
	"bookstore/infra/redis"
)

const (
	ServiceName = "Account Service"
)

func main() {
	db, err := database.Connect(
		os.Getenv("ACCOUNT_DB_USER"),
		os.Getenv("ACCOUNT_DB_PASSWORD"),
		os.Getenv("ACCOUNT_DB_HOST"),
		os.Getenv("ACCOUNT_DB_PORT"),
		os.Getenv("ACCOUNT_DB_DATABASE"),
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

	requestClient := request.NewHttpRequest(client.NewHTTPClient(), os.Getenv("CONSUMERS_HOST"))
	apiGatewayAuth := auth.NewAuth(requestClient, logger, os.Getenv("JWT_SECRET_KEY"))

	err = server.Start(dbService, cache, apiGatewayAuth, logger)
	if err != nil {
		log.Fatal(err.Error())
	}
}
