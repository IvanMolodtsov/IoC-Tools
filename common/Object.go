package common

type Object[T Any] interface {
	Getter[T]
	Setter[T]
	Remover
}
