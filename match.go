package match

import (
	"github.com/krelinga/go-match/matchfmt"
)

type Matcher[T any] interface {
	Match(got T) (matched bool, explanation string)
}

type MatcherFunc[T any] func(T) (bool, string)

func (f MatcherFunc[T]) Match(v T) (bool, string) {
	return f(v)
}

func Match[T any](got T, matcher Matcher[T]) (bool, string) {
	return matcher.Match(got)
}

func DefaultString[T any](v T) string {
	return matchfmt.DefaultFormat(v)
}
