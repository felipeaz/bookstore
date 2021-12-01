package pkg

import (
	"net/http"

	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/orders/order/model"
	"github.com/gin-gonic/gin"
)

// AssociateOrderInput is responsible for associating the params to the user model.
func AssociateOrderInput(c *gin.Context) (order model.Order, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&order)
	if err != nil {
		return model.Order{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
