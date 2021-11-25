package converter

import (
	"net/http"

	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/orders/items/model"
)

func ConvertToItemObj(obj interface{}) (model.Item, *errors.ApiError) {
	item, ok := obj.(*model.Item)
	if !ok {
		return model.Item{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *item, nil
}

func ConvertToSliceItemObj(obj interface{}) ([]model.Item, *errors.ApiError) {
	if obj == nil {
		return []model.Item{}, nil
	}
	items, ok := obj.(*[]model.Item)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *items, nil
}
