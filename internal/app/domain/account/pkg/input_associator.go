package pkg

import (
	"net/http"

	"bookstore/internal/app/domain/account/users/model"
	"github.com/gin-gonic/gin"

	"bookstore/internal/app/constants/errors"
)

// AssociateAccountInput is responsible for associating the params to the user model.
func AssociateAccountInput(c *gin.Context) (account model.Account, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&account)
	if err != nil {
		return model.Account{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}

// AssociateSessionInput is responsible for associating the params to the user model.
func AssociateSessionInput(c *gin.Context) (session model.UserSession, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&session)
	if err != nil {
		return model.UserSession{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
