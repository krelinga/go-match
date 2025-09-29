package pivot

import (
	"cmp"
	"fmt"
)

func NewOrderedValue[T cmp.Ordered](v T) OrderedValue[T] {
	return OrderedValue[T]{V: v}
}

type OrderedValue[T cmp.Ordered] struct {
	V T
}

func (ov OrderedValue[T]) String() string {
	return fmt.Sprintf("%#v", ov.V)
}

func (ov OrderedValue[T]) SameTypeAs(other any) bool {
	_, ok := other.(T)
	return ok
}

func (ov OrderedValue[T]) Compare(other any) int {
	return cmp.Compare(ov.V, other.(T))
}

func (ov OrderedValue[T]) Equal(other any) bool {
	return ov.Compare(other) == 0
}

func NewMapValue[K comparable, V any](m map[K]V) MapValue[K, V] {
	return MapValue[K, V]{M: m}
}

type MapValue[K comparable, V any] struct {
	M map[K]V
}

func (mv MapValue[K, V]) String() string {
	return fmt.Sprintf("%#v", mv.M)
}

func (mv MapValue[K, V]) SameTypeAs(other any) bool {
	_, ok := other.(map[K]V)
	return ok
}

func (mv MapValue[K, V]) Length() int {
	return len(mv.M)
}