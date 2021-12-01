package router

import (
	"bookstore/build/server/orders/router/build"
	"bookstore/internal/app/domain/orders/order/handler"
	"bookstore/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func Build(orderHandler handler.OrderHandler) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.OrderRoutes(vGroup, orderHandler)

	return router.Run(":8083")
}
