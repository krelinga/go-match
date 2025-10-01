package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
	"github.com/krelinga/go-typemap"
)

func notEqualImpl[T any](tm interface {
	typemap.String[T]
	typemap.Compare[T]
}, name string, other T) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = !tm.Compare(got, other)
		expected := fmt.Sprintf("got != %s", tm.String(other))
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

func NotEqualTm[T any](tm interface {
	typemap.String[T]
	typemap.Compare[T]
}, other T) Matcher[T] {
	return notEqualImpl(tm, "match.NotEqualTm", other)
}
