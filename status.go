package httpcodeerrors

import (
	"encoding/json"
)

type jsonMessage struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Message interface{} `json:"message"`
}

func (jm jsonMessage) stringify() string {
	b, err := json.Marshal(jm)
	if err != nil {
		panic(err)
	}
	return string(b)
}
