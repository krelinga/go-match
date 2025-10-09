package push

import "fmt"

type Formatter[T any] interface {
	Format(T) string
}

type defaultFormatter[T any] struct{}

func (f defaultFormatter[T]) Format(v T) string {
	return fmt.Sprintf("%#v", v)
}

func DefaultFormatter[T any]() Formatter[T] {
	return defaultFormatter[T]{}
}