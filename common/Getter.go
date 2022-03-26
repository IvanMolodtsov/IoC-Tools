package common

import (
	"fmt"
)

type Getter[T Any] interface {
	Get(string) (T, error)
}

type GetterError struct {
	key   string
	cause error
}

func (e GetterError) Error() string {
	return fmt.Sprintf("cannot get %s field: %s", e.key, e.cause.Error())
}
