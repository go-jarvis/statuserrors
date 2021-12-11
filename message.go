package statuserrors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func newMessage(code int, data interface{}) *message {

	m := &message{
		Code:    strconv.Itoa(code),
		Reason:  StatusText(code),
		Message: data,
	}

	return m
}

type message struct {
	Code    string      `json:"code"`
	Reason  string      `json:"reason"`
	Message interface{} `json:"message"`
}

func (jm *message) stringify() string {
	b, err := json.Marshal(jm)
	if err != nil {
		return fmt.Sprint(jm)
	}
	return string(b)
}

const (
	StatusUnknownError = 999
)

var statusText = map[int]string{
	999: "Unknown Error",
}

// StatusText returns a text for the HTTP status code.  Or text for status errors code.
func StatusText(code int) string {
	if s := http.StatusText(code); s != "" {
		return s
	}

	return statusText[code]
}
