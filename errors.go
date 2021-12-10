package httpcodeerrors

import (
	"net/http"
)

// New Status Error , http is http status code, data is message to reply
func New(httpcode int, data interface{}) error {
	e := &statusError{
		code: httpcode,
		data: data,
	}

	return e
}

// statusError define an error struct with httpcode
type statusError struct {
	err  error
	code int
	data interface{}
}

// Error return json format string
func (e *statusError) Error() string {
	jm := jsonMessage{
		Code:    e.code,
		Brief:   http.StatusText(e.code),
		Message: e.data,
	}

	return jm.stringify()

}

// Code return error code
func (e *statusError) Code() int {
	return e.code
}

// Data return error data
func (e *statusError) Data() interface{} {
	return e.data
}

type StatusError interface {
	HttpCode() int
	Data() interface{}
}
