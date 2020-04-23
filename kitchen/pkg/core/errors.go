package core

import (
	"fmt"
)

const (
	generalCat = (iota + 1) * 1000
)

var (
	ResourceNotFound    = newGeneralErr(1, "resource not found")
	InvalidOrderRequest = newGeneralErr(2, "invalid order")
)

type Error struct {
	ErrorCode int                    `json:"errorCode"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s",
		e.ErrorCode, e.Message)
}

// Is allows us to check the error with `errors.Is`
func (e *Error) Is(err error) bool {
	ae, ok := err.(*Error)
	if !ok {
		return false
	}
	return e.ErrorCode == ae.ErrorCode
}

func (e *Error) PutField(name string, value interface{}) {
	e.Fields[name] = value
}

// WithField creates an new MVError instance with the same code and message
// and set related field
func (e *Error) WithField(name string, value interface{}) *Error {
	err := new(Error)
	*err = *e
	err.Fields = make(map[string]interface{})
	err.PutField(name, value)
	return err
}

func newErr(code int, message string) *Error {
	return &Error{
		ErrorCode: code,
		Message:   message,
		Fields:    make(map[string]interface{}),
	}
}

func newGeneralErr(code int, message string) *Error {
	return newErr(code+generalCat, "general: "+message)
}
