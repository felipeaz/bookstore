package pkg

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/inventory/books/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAssociateBookInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expectBook := model.Book{
		ID:              1,
		Title:           "Clean Code",
		AuthorFirstName: "Robert",
		AuthorLastName:  "C Martin",
		Amount:          10,
	}
	b, _ := json.Marshal(expectBook)
	jsonParam := strings.NewReader(string(b))

	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	book, apiError := AssociateBookInput(c)

	assert.Nil(t, apiError)
	assert.Equal(t, expectBook.Title, book.Title)
	assert.Equal(t, expectBook.AuthorFirstName, book.AuthorFirstName)
	assert.Equal(t, expectBook.AuthorLastName, book.AuthorLastName)
	assert.Equal(t, expectBook.Amount, book.Amount)
	assert.Equal(t, expectBook.GetAuthorName(), book.GetAuthorName())
}

func TestAssociateBookInputWithError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(`{notABookProperty:"sample property"`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	book, apiError := AssociateBookInput(c)

	assert.Equal(t, model.Book{}, book)
	assert.Equal(t, http.StatusBadRequest, apiError.Status)
	assert.Equal(t, errors.FailedFieldsAssociationMessage, apiError.Message)
}
