package httpcodeerrors

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func newMessage(code int, data interface{}) *message {

	m := &message{
		Code:    strconv.Itoa(code),
		Reason:  http.StatusText(code),
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
		panic(err)
	}
	return string(b)
}
