package statuserrors

import "errors"

// Unwrap return error in httpcode errors
func (e *statusError) Unwrap() error {
	return e.err
}

// Wrap return a new errors with origin error and new error code state
func Wrap(err error, code int, data interface{}) error {
	return &statusError{
		code: code,
		data: data,
		err:  err,
	}
}

// Unwrap method is short for Unrap() in std libray
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Is method is short for Is() in std library
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As method is short for As() in std library
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
