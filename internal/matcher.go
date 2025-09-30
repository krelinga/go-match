package internal

// This is an exact copy of Matcher from the public API.  It is duplicated
// here to avoid an import cycle.
type Matcher[T any] interface {
	Match(v T) (matched bool, explanation string)
}

// This is an exact copy of MatcherFunc from the public API.  It is duplicated
// here to avoid an import cycle.
type MatcherFunc[T any] func(v T) (matched bool, explanation string)

func (f MatcherFunc[T]) Match(v T) (matched bool, explanation string) {
	return f(v)
}