package module

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/inventory/books/model"
	"bookstore/internal/app/domain/inventory/books/repository/interface"
	"bookstore/internal/app/domain/server"
	"bookstore/internal/app/logger"
	"context"
	_errors "errors"
	"net/http"
)

// BookModule process the request received from handler.
type BookModule struct {
	Repository _interface.BookRepositoryInterface
	Cache      database.CacheInterface
	Log        logger.LogInterface
}

func NewBookModule(
	repo _interface.BookRepositoryInterface,
	cache database.CacheInterface,
	log logger.LogInterface) BookModule {
	return BookModule{
		Repository: repo,
		Cache:      cache,
		Log:        log,
	}
}

// Get returns all books on DB.
func (m BookModule) Get() ([]model.Book, *errors.ApiError) {
	return m.Repository.Get()
}

// Find returns all books on DB.
func (m BookModule) Find(id string) (model.Book, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create persist a book to the database.
func (m BookModule) Create(book model.Book) (uint, *errors.ApiError) {
	return m.Repository.Create(book)
}

// Update update an existent book.
func (m BookModule) Update(id string, upBook model.Book) *errors.ApiError {
	return m.Repository.Update(id, upBook)
}

// Delete delete an existent book.
func (m BookModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}

// ChangeAmount receives an order and reduce the amount on inventory
func (m BookModule) ChangeAmount(ctx context.Context, req *server.Request) (*server.Response, error) {
	upBook, apiError := m.Find(req.BookId)
	if apiError != nil {
		return &server.Response{
			Success: false,
			Status:  int32(apiError.Status),
		}, apiError.GetError()
	}

	if upBook.Amount == 0 {
		return &server.Response{
			Success: false,
			Status:  http.StatusBadRequest,
		}, _errors.New("item is out of stock")
	}

	updateAmount := upBook.Amount - req.Amount
	if updateAmount < 0 {
		return &server.Response{
			Success: false,
			Status:  http.StatusBadRequest,
		}, _errors.New("amount requested is greater than available stock")
	}
	upBook.Amount = updateAmount

	apiError = m.Update(req.BookId, upBook)
	if apiError != nil {
		return &server.Response{
			Success: false,
			Status:  int32(apiError.Status),
		}, apiError.GetError()
	}

	return &server.Response{Success: true}, nil
}
