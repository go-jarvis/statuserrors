package statuserrors

import (
	stderrors "errors"
	"fmt"
	"net/http"
	"testing"
)

func Test_JsonStringify(t *testing.T) {
	data := struct {
		Name string
		Age  int `json:"int"`
	}{
		Name: "zhangsan",
		Age:  10,
	}

	jm := newMessage(http.StatusOK, data)

	fmt.Println(jm)
	s := jm.stringify()
	fmt.Println(s)
}

func Test_ErrorWrap(t *testing.T) {
	err := newError(200, "success")
	fmt.Println(err)

	err = Wrap(err, http.StatusNotFound, "balalalla")
	fmt.Println(err)

	err = stderrors.Unwrap(err)
	fmt.Println(err)

	// pkgerr := pkgerrors.New("123")
	// pkgerrors.Wrap()
}

func newError(code int, str string) error {
	return New(http.StatusOK, str)
}
