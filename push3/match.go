package push3

func Match[T any](in T, m M[T]) Result {
	return m(in)
}