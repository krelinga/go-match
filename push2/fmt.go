package push2

import "fmt"

type FmtFunc[T any] func(T) string

func DefaultFmtFunc[T any]() FmtFunc[T] {
	return func(v T) string {
		return fmt.Sprintf("%#v", v)
	}
}