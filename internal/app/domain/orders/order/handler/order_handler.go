package handler

import (
	"net/http"

	_interface "bookstore/internal/app/domain/orders/order/module/interface"
	"bookstore/internal/app/domain/orders/pkg"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Module _interface.OrderModuleInterface
}

func NewOrderHandler(module _interface.OrderModuleInterface) OrderHandler {
	return OrderHandler{
		Module: module,
	}
}

func (h OrderHandler) Get(c *gin.Context) {
	orders, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h OrderHandler) Find(c *gin.Context) {
	order, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func (h OrderHandler) Create(c *gin.Context) {
	order, apiError := pkg.AssociateOrderInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	id, apiError := h.Module.Create(order)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h OrderHandler) Update(c *gin.Context) {
	upOrder, apiError := pkg.AssociateOrderInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upOrder)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h OrderHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}
