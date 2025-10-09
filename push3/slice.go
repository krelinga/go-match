package push3

type Slice[T ~[]E, E any] struct {
	raw T
}

func (s Slice[T, E]) Len() int {
	return len(s.raw)
}

func (s Slice[T, E]) Nil() bool {
	return s.raw == nil
}

func WrapSlice[T ~[]E, E any](t T) Slice[T, E] {
	return Slice[T, E]{raw: t}
}