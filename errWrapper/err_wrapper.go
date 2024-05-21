package errWrapper

import (
	"fmt"
	"github.com/pkg/errors"
)

type (
	ErrorType uint
)
type customError struct {
	errorType     ErrorType
	originalError error
	context       errorContext
}
type errorContext struct {
	Field   string
	Message string
}

const (
	NoType ErrorType = iota
	BadRequest
	NotFound
)

func (errorType ErrorType) New(msg string) error {
	return customError{
		errorType:     errorType,
		originalError: errors.New(msg),
	}
}
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{
		errorType:     errorType,
		originalError: errors.Wrapf(err, msg, args...),
	}
}

func New(msg string) error {
	return customError{errorType: NoType, originalError: errors.New(msg)}
}
func NewF(msg string, arg ...interface{}) error {
	return customError{errorType: NoType, originalError: errors.New(fmt.Sprintf(msg, arg...))}
}
func (error customError) Error() string {
	return error.originalError.Error()
}
func AddErrorContext(err error, field, msg string) error {
	context := errorContext{Field: field, Message: msg}
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: customErr.originalError,
			context:       context,
		}
	}

	return customError{errorType: NoType, originalError: err, context: context}
}
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}

	if customErr, ok := err.(customError); ok || customErr.context != emptyContext {
		return map[string]string{
			"field":   customErr.context.Field,
			"massage": customErr.context.Message,
		}
	}

	return nil
}

func Wrapf(err error, msg string, args ...interface{}) error {
	wrapperError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrapperError,
			context:       customErr.context,
		}
	}

	return customError{
		errorType:     NoType,
		originalError: wrapperError,
	}
}
