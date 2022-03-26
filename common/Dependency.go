package common

import "fmt"

type DependencyError struct {
	cause error
}

func (e DependencyError) Error() string {
	return fmt.Sprintf("Dependency invokation failed: %s", e.cause.Error())
}

type Dependency func([]Any) (Any, error)
