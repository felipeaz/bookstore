package pkg

//import (
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//
//	bookModel "bookstore/internal/app/domain/inventory/books/model"
//	categoryModel "bookstore/internal/app/domain/inventory/categories/model"
//	lendingModel "bookstore/internal/app/domain/inventory/lending/model"
//	studentModel "bookstore/internal/app/domain/inventory/students/model"
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestAssociateBookInput(t *testing.T) {
//	// Init
//	gin.SetMode(gin.TestMode)
//	jsonParam := strings.NewReader(`{"title":"Code Universe","author":"Unknown Author","registerNumber":"123456"}`)
//	req := http.Request{Body: ioutil.NopCloser(jsonParam)}
//
//	w := httptest.NewRecorder()
//	c, _ := gin.CreateTestContext(w)
//	c.Request = &req
//
//	// Execution
//	book, apiError := AssociateBookInput(c)
//
//	// Validation
//	assert.Nil(t, apiError)
//	assert.Equal(t, "Code Universe", book.Title)
//	assert.Equal(t, "Unknown Author", book.Author)
//	assert.Equal(t, "123456", book.RegisterNumber)
//}
//
//func TestAssociateBookInputWithError(t *testing.T) {
//	// Init
//	gin.SetMode(gin.TestMode)
//	jsonParam := strings.NewReader(`{"title":"Code Universe","author":"Unknown Author","registerNumber":""}`)
//	req := http.Request{Body: ioutil.NopCloser(jsonParam)}
//
//	w := httptest.NewRecorder()
//	c, _ := gin.CreateTestContext(w)
//	c.Request = &req
//
//	// Execution
//	book, apiError := AssociateBookInput(c)
//
//	// Validation
//	assert.Equal(t, bookModel.Book{}, book)
//	assert.Equal(t, http.StatusBadRequest, apiError.Status)
//}
