package common

type Appendable[T Any] interface {
	Append(T)
}
