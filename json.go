package statuserrors

import (
	"encoding/json"
	"fmt"
)

type jsonMessage struct {
	Code    int         `json:"code"`
	Brief   string      `json:"brief"`
	Message interface{} `json:"message"`
}

func (jm *jsonMessage) stringify() string {
	fmt.Println(jm)
	// b, _ := json.MarshalIndent(&jm, "", "")
	b, err := json.Marshal(jm)
	if err != nil {
		panic(err)
	}
	return string(b)
}
