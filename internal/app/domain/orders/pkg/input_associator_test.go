package pkg

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/domain/orders/order/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAssociateOrderInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expectOrder := model.Order{
		ID:              1,
		BookId:          2,
		ClientFirstName: "Felipe",
		ClientLastName:  "Silva",
		Amount:          10,
	}
	b, _ := json.Marshal(expectOrder)
	jsonParam := strings.NewReader(string(b))

	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	order, apiError := AssociateOrderInput(c)

	assert.Nil(t, apiError)
	assert.Equal(t, expectOrder.ID, order.ID)
	assert.Equal(t, expectOrder.BookId, order.BookId)
	assert.Equal(t, expectOrder.ClientFirstName, order.ClientFirstName)
	assert.Equal(t, expectOrder.ClientLastName, order.ClientLastName)
	assert.Equal(t, expectOrder.Amount, order.Amount)
	assert.Equal(t, expectOrder.GetClientName(), order.GetClientName())
}

func TestAssociateOrderInputWithError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(`{notAOrderProperty:"sample property"`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	order, apiError := AssociateOrderInput(c)

	assert.Equal(t, model.Order{}, order)
	assert.Equal(t, http.StatusBadRequest, apiError.Status)
	assert.Equal(t, errors.FailedFieldsAssociationMessage, apiError.Message)
}
