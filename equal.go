package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
	"github.com/krelinga/go-typemap"
)

func equalImpl[T any](tm interface {
	typemap.String[T]
	typemap.Compare[T]
}, name string, want T) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = tm.Compare(got, want)
		expected := fmt.Sprintf("got == %s", tm.String(want))
		var info string
		if !matched {
			actual := fmt.Sprintf("got == %s", tm.String(got))
			info = matchutil.ActualVsExpected(actual, expected)
		} else {
			info = expected
		}
		explanation = matchutil.Explain(matched, name, info)
		return
	})
}

func EqualTm[T any](tm interface {
	typemap.String[T]
	typemap.Compare[T]
}, want T) Matcher[T] {
	return equalImpl(tm, "match.EqualTm", want)
}
