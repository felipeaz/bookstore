package handler

import (
	"net/http"

	_interface "bookstore/internal/app/domain/orders/items/module/interface"
	"bookstore/internal/app/domain/orders/pkg"
	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	Module _interface.ItemModuleInterface
}

func NewItemHandler(module _interface.ItemModuleInterface) ItemHandler {
	return ItemHandler{
		Module: module,
	}
}

func (h ItemHandler) Get(c *gin.Context) {
	items, apiError := h.Module.Get(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h ItemHandler) Find(c *gin.Context) {
	item, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}

func (h ItemHandler) Create(c *gin.Context) {
	item, apiError := pkg.AssociateItemInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	id, apiError := h.Module.Create(item)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h ItemHandler) Update(c *gin.Context) {
	upItem, apiError := pkg.AssociateItemInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upItem)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h ItemHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}
