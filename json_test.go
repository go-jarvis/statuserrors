package httpcodeerrors

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_JsonStringify(t *testing.T) {
	jm := jsonMessage{
		Code:  200,
		Brief: http.StatusText(200),
		Message: struct {
			Name string
			Age  int `json:"int"`
		}{
			Name: "zhangsan",
			Age:  10,
		},
	}

	fmt.Println(jm)
	s := jm.stringify()
	fmt.Println(s)
}
