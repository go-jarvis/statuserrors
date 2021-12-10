package httpcodeerrors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func Test_JsonStringify(t *testing.T) {
	jm := jsonMessage{
		Code:   200,
		Reason: http.StatusText(200),
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

func Test_ErrorWrap(t *testing.T) {
	err := newError(200, "success")
	fmt.Println(err)

	err = Wrap(http.StatusNotFound, "balalalla", err)
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)

}

func newError(code int, str string) error {
	return New(http.StatusOK, str)
}
