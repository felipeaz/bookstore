package _interface

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/inventory/books/model"
	"bookstore/internal/app/domain/server"
	"context"
)

type BookModuleInterface interface {
	Get() ([]model.Book, *errors.ApiError)
	Find(id string) (model.Book, *errors.ApiError)
	Create(book model.Book) (model.Book, *errors.ApiError)
	Update(id string, upBook model.Book) *errors.ApiError
	Delete(id string) *errors.ApiError
	ChangeAmount(ctx context.Context, req *server.Request) (*server.Response, error)
}
