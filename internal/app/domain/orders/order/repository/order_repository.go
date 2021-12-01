package repository

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/orders/order/model"
	"bookstore/internal/app/domain/orders/order/model/converter"
)

type OrderRepository struct {
	DB database.GORMServiceInterface
}

func NewOrderRepository(db database.GORMServiceInterface) OrderRepository {
	return OrderRepository{
		DB: db,
	}
}

func (r OrderRepository) Get() ([]model.Order, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]model.Order{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	orders, apiError := converter.ConvertToSliceOrderObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return orders, nil
}

func (r OrderRepository) Find(id string) (model.Order, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Order{}, id)
	if err != nil {
		return model.Order{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	order, apiError := converter.ConvertToOrderObj(result)
	if apiError != nil {
		return model.Order{}, apiError
	}
	return order, nil
}

func (r OrderRepository) Create(order model.Order) (model.Order, *errors.ApiError) {
	err := r.DB.Persist(&order)
	if err != nil {
		return model.Order{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return order, nil
}

func (r OrderRepository) Update(id string, upOrder model.Order) *errors.ApiError {
	err := r.DB.Refresh(&upOrder, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r OrderRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&model.Order{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
