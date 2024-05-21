package errWrapper

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext(t *testing.T) {
	err := BadRequest.New("error")
	errWithContext := AddErrorContext(err, "field 2", "field is zero")
	expectedContext := map[string]string{"field": "field 2", "massage": "field is zero"}

	assert.Equal(t, BadRequest, GetType(errWithContext))
	assert.Equal(t, NoType, GetType(errors.New("123")))
	assert.Equal(t, err.Error(), errWithContext.Error())
	assert.Equal(t, expectedContext, GetErrorContext(errWithContext))
}

func TestWrapf(t *testing.T) {
	err := New("error")
	wrapperError := BadRequest.Wrapf(err, "error %s", "1")

	assert.Equal(t, BadRequest, GetType(wrapperError))
	assert.EqualError(t, wrapperError, "error 1: error")
}

func TestWrapfIfNotTypeError(t *testing.T) {
	err := NewF("error %s", "2")
	wrapperError := Wrapf(err, "error %s", "1")

	assert.Equal(t, NoType, GetType(wrapperError))
	assert.EqualError(t, wrapperError, "error 1: error 2")
}
