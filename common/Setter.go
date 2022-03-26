package common

import (
	"fmt"
)

type Setter[T Any] interface {
	Set(string, T) error
}

type SetterError struct {
	key   string
	value Any
	cause error
}

func (e SetterError) Error() string {
	return fmt.Sprintf("cannot set %s field with %v: %s", e.key, e.value, e.cause.Error())
}
