package statuserrors

import (
	stderr "errors"
	"fmt"
)

// Unwrap return error in httpcode errors
func (e *statusError) Unwrap() error {
	return e.err
}

// Wrap return a new errors with origin error and new error code state
func Wrap(err error, code int, format string, a ...interface{}) *statusError {
	return &statusError{
		code:    code,
		message: fmt.Sprintf(format, a...),
		err:     err,
	}
}

// Unwrap method is short for Unrap() in std libray
func Unwrap(err error) error {
	return stderr.Unwrap(err)
}

// Is method is short for Is() in std library
func Is(err, target error) bool {
	return stderr.Is(err, target)
}

// As method is short for As() in std library
func As(err error, target interface{}) bool {
	return stderr.As(err, target)
}
