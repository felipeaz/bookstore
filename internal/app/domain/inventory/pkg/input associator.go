package pkg

import (
	"net/http"

	"bookstore/internal/app/constants/errors"
	bookModel "bookstore/internal/app/domain/inventory/books/model"
	"github.com/gin-gonic/gin"
)

// AssociateBookInput is responsible for associate the params to the book model.
func AssociateBookInput(c *gin.Context) (book bookModel.Book, apiError *errors.ApiError) {
	err := c.ShouldBindJSON(&book)
	if err != nil {
		return bookModel.Book{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedFieldsAssociationMessage,
			Error:   err.Error(),
		}
	}

	return
}
