package push

type Result struct {
	Matched bool
	Reason  string
}

type M[T any] func(T) Result

func Match[T any](got T, matcher M[T]) Result {
	return matcher(got)
}