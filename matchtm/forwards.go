package matchtm

import (
	"github.com/krelinga/go-match"
	"github.com/krelinga/go-match/internal"
	"github.com/krelinga/go-typemap"
)

func Length[T any](tm typemap.Length[T], lengthMatcher match.Matcher[int]) match.Matcher[T] {
	return internal.Length(tm, lengthMatcher)
}