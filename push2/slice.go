package push2

type Slice[T ~[]E, E any] struct {
	raw T
}

func (s Slice[T, E]) Len() int {
	return len(s.raw)
}

func (s Slice[T, E]) IsNil() bool {
	return s.raw == nil
}

func AsSlice[T ~[]E, E any](matcher Matcher[Slice[T, E]]) Matcher[T] {
	return Func[T](func(t T) Result {
		return matcher.Match(Slice[T, E]{raw: t})
	})
}

func WrapSlice[T ~[]E, E any](t T) Slice[T, E] {
	return Slice[T, E]{raw: t}
}