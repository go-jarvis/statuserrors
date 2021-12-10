package statuserrors

// interface checking
var _ StatusError = (*statusError)(nil)

type StatusError interface {
	StatusCode() int
}

// New Status Error , code is http status code or status errors code, message is error reason
func New(code int, message string) *statusError {
	e := &statusError{
		code:    code,
		message: message,
	}

	return e
}

// statusError define an error struct with httpcode
type statusError struct {
	err     error
	code    int
	message string
}

// Error return json format string
func (e *statusError) Error() string {
	msg := newMessage(e.code, e.message)

	return msg.stringify()
}

// Code return error code
func (e *statusError) StatusCode() int {
	return e.code
}
