package _interface

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/inventory/books/model"
)

type BookModuleInterface interface {
	Get() ([]model.Book, *errors.ApiError)
	Find(id string) (model.Book, *errors.ApiError)
	Create(book model.Book) (uint, *errors.ApiError)
	Update(id string, upBook model.Book) *errors.ApiError
	Delete(id string) *errors.ApiError
}
