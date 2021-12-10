package statuserrors

// interface checking
var _ StatusError = (*statusError)(nil)

type StatusError interface {
	Code() int
	Data() interface{}
}

// New Status Error , http is http status code, message is error reason
func New(httpcode int, message string) error {
	e := &statusError{
		code:    httpcode,
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
func (e *statusError) Code() int {
	return e.code
}

// Data return error data
func (e *statusError) Data() interface{} {
	return e.message
}
