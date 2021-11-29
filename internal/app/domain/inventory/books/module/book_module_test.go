package module

import (
	"bookstore/infra/logger"
	mysql "bookstore/infra/mysql/service"
	"bookstore/infra/redis"
	"bookstore/internal/app/domain/inventory/books/constants"
	bookModel "bookstore/internal/app/domain/inventory/books/model"
	"bookstore/internal/app/domain/inventory/books/repository"
	"bookstore/internal/app/domain/server"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestChangeAmount(t *testing.T) {
	dbMock := new(mysql.MockMySQLService)
	cacheMock := new(redis.MockCache)
	logMock := new(logger.Mock)

	req := &server.Request{
		BookId: "355",
		Amount: 5,
	}
	bookObj := bookModel.Book{
		ID:              355,
		Title:           "Clean Code",
		AuthorFirstName: "Robert",
		AuthorLastName:  "C. Martin",
		Amount:          15,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	upBook := bookObj
	upBook.Amount = bookObj.Amount - req.Amount
	b, _ := json.Marshal(bookObj)

	repo := repository.NewBookRepository(dbMock)
	service := NewBookModule(repo, cacheMock, logMock)

	dbMock.On("find", req.BookId).Return(bookObj, nil).Once()
	dbMock.On("Refresh", req.BookId).Return(nil).Once()
	cacheMock.On("Get", req.BookId).Return(b, nil).Once()
	cacheMock.On("Flush", req.BookId).Return(nil).Once()
	cacheMock.On("Flush", "books").Return(nil).Once()

	resp, err := service.ChangeAmount(context.Background(), req)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Success, true)
}

func TestChangeAmountRequestedGreaterThanStock(t *testing.T) {
	dbMock := new(mysql.MockMySQLService)
	cacheMock := new(redis.MockCache)
	logMock := new(logger.Mock)

	req := &server.Request{
		BookId: "355",
		Amount: 5,
	}
	bookObj := bookModel.Book{
		ID:              355,
		Title:           "Clean Code",
		AuthorFirstName: "Robert",
		AuthorLastName:  "C. Martin",
		Amount:          2,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	upBook := bookObj
	upBook.Amount = bookObj.Amount - req.Amount
	b, _ := json.Marshal(bookObj)

	repo := repository.NewBookRepository(dbMock)
	service := NewBookModule(repo, cacheMock, logMock)

	dbMock.On("find", req.BookId).Return(bookObj, nil).Once()
	dbMock.On("Refresh", req.BookId).Return(nil).Once()
	cacheMock.On("Get", req.BookId).Return(b, nil).Once()
	cacheMock.On("Flush", req.BookId).Return(nil).Once()
	cacheMock.On("Flush", "books").Return(nil).Once()

	resp, err := service.ChangeAmount(context.Background(), req)

	assert.NotNil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Status, int32(http.StatusBadRequest))
	assert.Equal(t, err.Error(), constants.AmountGreaterThanStockError)
	assert.Equal(t, resp.Success, false)
}

func TestChangeAmountItemOutOfStock(t *testing.T) {
	dbMock := new(mysql.MockMySQLService)
	cacheMock := new(redis.MockCache)
	logMock := new(logger.Mock)

	req := &server.Request{
		BookId: "355",
		Amount: 5,
	}
	bookObj := bookModel.Book{
		ID:              355,
		Title:           "Clean Code",
		AuthorFirstName: "Robert",
		AuthorLastName:  "C. Martin",
		Amount:          0,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	upBook := bookObj
	upBook.Amount = bookObj.Amount - req.Amount
	b, _ := json.Marshal(bookObj)

	repo := repository.NewBookRepository(dbMock)
	service := NewBookModule(repo, cacheMock, logMock)

	dbMock.On("find", req.BookId).Return(bookObj, nil).Once()
	dbMock.On("Refresh", req.BookId).Return(nil).Once()
	cacheMock.On("Get", req.BookId).Return(b, nil).Once()
	cacheMock.On("Flush", req.BookId).Return(nil).Once()
	cacheMock.On("Flush", "books").Return(nil).Once()

	resp, err := service.ChangeAmount(context.Background(), req)

	assert.NotNil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Status, int32(http.StatusBadRequest))
	assert.Equal(t, err.Error(), constants.ItemOutOfStockError)
	assert.Equal(t, resp.Success, false)
}
