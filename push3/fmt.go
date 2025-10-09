package push3

import "fmt"

type Fmt[T any] func(T) string

func DefaultFmt[T any]() Fmt[T] {
	return func(v T) string {
		return fmt.Sprintf("%#v", v)
	}
}