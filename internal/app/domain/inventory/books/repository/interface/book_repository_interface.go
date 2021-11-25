package _interface

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/inventory/books/model"
)

type BookRepositoryInterface interface {
	Get() (books []model.Book, apiError *errors.ApiError)
	Find(id string) (book model.Book, apiError *errors.ApiError)
	Create(book model.Book) (uint, *errors.ApiError)
	Update(id string, upBook model.Book) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
