package converter

import (
	"encoding/json"
	"net/http"

	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/orders/order/model"
)

func ConvertToOrderObj(obj interface{}) (model.Order, *errors.ApiError) {
	order, ok := obj.(*model.Order)
	if !ok {
		return model.Order{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *order, nil
}

func ConvertToSliceOrderObj(obj interface{}) ([]model.Order, *errors.ApiError) {
	if obj == nil {
		return []model.Order{}, nil
	}
	orders, ok := obj.(*[]model.Order)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *orders, nil
}

func ConvertToOrderObjFromCache(b []byte) (model.Order, *errors.ApiError) {
	var order model.Order
	err := json.Unmarshal(b, &order)
	if err != nil {
		return model.Order{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
			Error:   err.Error(),
		}
	}
	return order, nil
}

func ConvertToSliceOrderObjFromCache(b []byte) ([]model.Order, *errors.ApiError) {
	var orders []model.Order
	err := json.Unmarshal(b, &orders)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
			Error:   err.Error(),
		}
	}
	return orders, nil
}
