package pkg

import (
	"bookstore/internal/app/constants/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIdParserSuccess(t *testing.T) {
	var expect uint
	expect = 355
	input := "355"

	resp, err := ParseStringToId(input)

	assert.Empty(t, err)
	assert.Equal(t, expect, resp)
}

func TestIdParserError(t *testing.T) {
	var expect uint
	expect = 0
	input := "abcd"

	resp, err := ParseStringToId(input)

	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, errors.FailedFieldsAssociationMessage, err.Message)
	assert.Equal(t, expect, resp)
}
