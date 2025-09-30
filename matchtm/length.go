package matchtm

import (
	"github.com/krelinga/go-match"
	"github.com/krelinga/go-match/matchutil"
	"github.com/krelinga/go-typemap"
)

func Length[T any](matcher match.Matcher[int], ht typemap.Length[T]) match.Matcher[T] {
	return match.MatcherFunc[T](func(v T) (matched bool, explanation string) {
		length := ht.Length(v)
		matched, e := matcher.Match(length)
		explanation = matchutil.Explain(matched, "matchtm.Length", e)
		return
	})
}
