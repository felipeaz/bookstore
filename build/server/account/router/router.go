package router

import (
	"bookstore/build/server/account/router/build"
	"bookstore/internal/app/domain/account/users/handler"
	"bookstore/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func Build(accountHandler handler.AccountHandler) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.UserRoutes(vGroup, accountHandler)

	return router.Run(":8081")
}
