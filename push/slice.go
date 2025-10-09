package push

type Element[KT, VT any] struct {
	Key KT
	Val VT
}

type Slice[T ~[]E, E any] struct {
	raw T
}

func (s Slice[T, E]) Len() int {
	return len(s.raw)
}

func (s Slice[T, E]) IsNil() bool {
	return s.raw == nil
}

func (s Slice[T, E]) Elements() []Element[int, E] {
	elems := make([]Element[int, E], 0, s.Len())
	for i, t := range s.raw {
		elems = append(elems, Element[int, E]{Key: i, Val: t})
	}
	return elems
}

func AsSlice[T ~[]E, E any](t T) Slice[T, E] {
	return Slice[T, E]{raw: t}
}
