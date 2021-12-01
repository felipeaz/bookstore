package _interface

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/orders/order/model"
)

type OrderRepositoryInterface interface {
	Get() ([]model.Order, *errors.ApiError)
	Find(id string) (model.Order, *errors.ApiError)
	Create(order model.Order) (model.Order, *errors.ApiError)
	Update(id string, upOrder model.Order) *errors.ApiError
	Delete(id string) *errors.ApiError
}
