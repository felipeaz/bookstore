package _interface

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/orders/items/model"
)

type ItemModuleInterface interface {
	Get() ([]model.Item, *errors.ApiError)
	Find(id string) (model.Item, *errors.ApiError)
	Create(item model.Item) (uint, *errors.ApiError)
	Update(id string, upItem model.Item) *errors.ApiError
	Delete(id string) *errors.ApiError
}
