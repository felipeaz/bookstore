package router

import (
	"bookstore/build/server/orders/router/build"
	"bookstore/internal/app/domain/orders/items/handler"
	"bookstore/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func Build(itemHandler handler.ItemHandler) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.ItemRoutes(vGroup, itemHandler)

	return router.Run(":8083")
}
