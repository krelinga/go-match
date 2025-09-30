package internal

import (
	"github.com/krelinga/go-match/matchutil"
	"github.com/krelinga/go-typemap"
)

func Length[T any](tm typemap.Length[T], matcher Matcher[int]) Matcher[T] {
	return MatcherFunc[T](func(v T) (matched bool, explanation string) {
		length := tm.Length(v)
		matched, e := matcher.Match(length)
		explanation = matchutil.Explain(matched, "matchtm.Length", e)
		return
	})
}
