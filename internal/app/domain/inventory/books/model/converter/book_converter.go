package converter

import (
	"net/http"

	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/inventory/books/model"
)

func ConvertToBookObj(obj interface{}) (model.Book, *errors.ApiError) {
	bookObj, ok := obj.(*model.Book)
	if !ok {
		return model.Book{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *bookObj, nil
}

func ConvertToSliceBookObj(obj interface{}) ([]model.Book, *errors.ApiError) {
	if obj == nil {
		return []model.Book{}, nil
	}
	bookObj, ok := obj.(*[]model.Book)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *bookObj, nil
}
