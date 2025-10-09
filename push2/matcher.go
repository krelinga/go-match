package push2

type Matcher[T any] interface {
	Match(T) Result
}

type Func[T any] func(T) Result

func (f Func[T]) Match(v T) Result {
	return f(v)
}

func Match[T any](got T, matcher Matcher[T]) Result {
	return matcher.Match(got)
}