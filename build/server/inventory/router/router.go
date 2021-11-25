package router

import (
	"bookstore/build/server/inventory/router/build"
	"bookstore/internal/app/domain/inventory/books/handler"
	"bookstore/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func Build(bookHandler handler.BookHandler) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.BookRoutes(vGroup, bookHandler)

	return router.Run(":8082")
}
