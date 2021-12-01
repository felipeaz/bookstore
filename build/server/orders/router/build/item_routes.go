package build

import (
	"bookstore/internal/app/domain/orders/order/handler"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(rg *gin.RouterGroup, orderHandler handler.OrderHandler) {
	r := rg.Group("/orders")
	r.GET("/", orderHandler.Get)
	r.GET("/:id", orderHandler.Find)
	r.POST("/", orderHandler.Create)
	r.PUT("/:id", orderHandler.Update)
	r.DELETE("/:id", orderHandler.Delete)
}
