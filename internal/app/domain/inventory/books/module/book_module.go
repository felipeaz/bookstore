package module

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/inventory/books/model"
	"bookstore/internal/app/domain/inventory/books/model/converter"
	"bookstore/internal/app/domain/inventory/books/repository/interface"
	"bookstore/internal/app/domain/server"
	"bookstore/internal/app/logger"
	"context"
	"encoding/json"
	_errors "github.com/pkg/errors"
	"net/http"
	"strconv"
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

const (
	AllData = "books"
)

func (m BookModule) Get() ([]model.Book, *errors.ApiError) {
	b, err := m.Cache.Get(AllData)
	if err != nil {
		m.Log.Error(err)
	}
	if b != nil {
		return converter.ConvertToSliceBookObjFromCache(b)
	}

	books, apiError := m.Repository.Get()
	if apiError != nil {
		return nil, apiError
	}
	m.setAllBookCache(books)
	return m.Repository.Get()
}

func (m BookModule) Find(id string) (model.Book, *errors.ApiError) {
	b, err := m.Cache.Get(id)
	if err != nil {
		m.Log.Error(err)
	}
	if b != nil {
		return converter.ConvertToBookObjFromCache(b)
	}

	book, apiError := m.Repository.Find(id)
	if apiError != nil {
		return model.Book{}, apiError
	}
	m.setBookCache(book)
	return book, nil
}

func (m BookModule) Create(book model.Book) (model.Book, *errors.ApiError) {
	book, apiError := m.Repository.Create(book)
	if apiError != nil {
		return model.Book{}, apiError
	}

	err := m.Cache.Flush(AllData)
	if err != nil {
		m.Log.Error(err)
	}
	m.setBookCache(book)
	return book, nil
}

func (m BookModule) Update(id string, upBook model.Book) *errors.ApiError {
	apiError := m.Repository.Update(id, upBook)
	if apiError != nil {
		return apiError
	}

	err := m.Cache.Flush(id)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Flush(AllData)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToFlushAllCache)
		m.Log.Error(err)
	}
	return nil
}

func (m BookModule) Delete(id string) *errors.ApiError {
	apiError := m.Repository.Delete(id)
	if apiError != nil {
		return apiError
	}

	err := m.Cache.Flush(id)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Flush(AllData)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToFlushAllCache)
		m.Log.Error(err)
	}
	return nil
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

func (m BookModule) setBookCache(book model.Book) {
	b, err := json.Marshal(book)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Set(strconv.FormatUint(uint64(book.ID), 10), b)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToSetCache)
	}
}

func (m BookModule) setAllBookCache(books []model.Book) {
	b, err := json.Marshal(books)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Set(AllData, b)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToSetCache)
	}
}
