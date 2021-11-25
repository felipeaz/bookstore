package converter

import (
	"net/http"

	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/account/users/model"
)

func ConvertToAccountObj(obj interface{}) (model.Account, *errors.ApiError) {
	accountObj, ok := obj.(*model.Account)
	if !ok {
		return model.Account{}, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *accountObj, nil
}

func ConvertToSliceAccountObj(obj interface{}) ([]model.Account, *errors.ApiError) {
	if obj == nil {
		return []model.Account{}, nil
	}
	accountObj, ok := obj.(*[]model.Account)
	if !ok {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToConvertObj,
		}
	}
	return *accountObj, nil
}
