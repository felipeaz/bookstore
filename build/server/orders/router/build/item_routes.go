package build

import (
	"bookstore/internal/app/domain/orders/items/handler"
	"github.com/gin-gonic/gin"
)

func ItemRoutes(rg *gin.RouterGroup, itemHandler handler.ItemHandler) {
	r := rg.Group("/orders")
	r.GET("/:id/book", itemHandler.Get)
	r.GET("/:id", itemHandler.Find)
	r.POST("/", itemHandler.Create)
	r.PUT("/:id", itemHandler.Update)
	r.DELETE("/:id", itemHandler.Delete)
}
