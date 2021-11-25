package pkg

import (
	"net/http"

	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/orders/items/model"
	"github.com/gin-gonic/gin"
)

// AssociateItemInput is responsible for associating the params to the user model.
func AssociateItemInput(c *gin.Context) (item model.Item, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&item)
	if err != nil {
		return model.Item{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
