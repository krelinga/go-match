package push2

func MatchLen(len int, matcher Matcher[int]) Result {
	return matcher.Match(len)
}

type ILen interface {
	Len() int
}

func AsILen[T ILen](m Matcher[ILen]) Matcher[T] {
	return Func[T](func(v T) Result {
		return m.Match(v)
	})
}

func Len(m Matcher[int]) Matcher[ILen] {
	return Func[ILen](func(v ILen) Result {
		return m.Match(v.Len())
	})
}