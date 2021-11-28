package converter

import (
	"encoding/json"
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

func ConvertToBookObjFromCache(b []byte) (model.Book, *errors.ApiError) {
	var book model.Book
	err := json.Unmarshal(b, &book)
	if err != nil {
		return model.Book{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
			Error:   err.Error(),
		}
	}
	return book, nil
}

func ConvertToSliceBookObjFromCache(b []byte) ([]model.Book, *errors.ApiError) {
	var books []model.Book
	err := json.Unmarshal(b, &books)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
			Error:   err.Error(),
		}
	}
	return books, nil
}
