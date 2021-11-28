package pkg

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/orders/items/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAssociateItemInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expectItem := model.Item{
		ID:              1,
		BookId:          2,
		ClientFirstName: "Felipe",
		ClientLastName:  "Silva",
		Amount:          10,
	}
	b, _ := json.Marshal(expectItem)
	jsonParam := strings.NewReader(string(b))

	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	item, apiError := AssociateItemInput(c)

	assert.Nil(t, apiError)
	assert.Equal(t, expectItem.ID, item.ID)
	assert.Equal(t, expectItem.BookId, item.BookId)
	assert.Equal(t, expectItem.ClientFirstName, item.ClientFirstName)
	assert.Equal(t, expectItem.ClientLastName, item.ClientLastName)
	assert.Equal(t, expectItem.Amount, item.Amount)
	assert.Equal(t, expectItem.GetClientName(), item.GetClientName())
}

func TestAssociateItemInputWithError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(`{notAItemProperty:"sample property"`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	item, apiError := AssociateItemInput(c)

	assert.Equal(t, model.Item{}, item)
	assert.Equal(t, http.StatusBadRequest, apiError.Status)
	assert.Equal(t, errors.FailedFieldsAssociationMessage, apiError.Message)
}
