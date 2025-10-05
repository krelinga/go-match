package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchfmt"
	"github.com/krelinga/go-typemap"
)

func lessThanImpl[T any](tm interface {
	typemap.String[T]
	typemap.Order[T]
}, name string, other T) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = tm.Order(got, other) < 0
		expected := fmt.Sprintf("got < %s", tm.String(other))
		var info string
		if !matched {
			actual := fmt.Sprintf("got == %s", tm.String(got))
			info = matchfmt.ActualVsExpected(actual, expected)
		} else {
			info = expected
		}
		explanation = matchfmt.Explain(matched, name, info)
		return
	})
}

func LessThanTm[T any](tm interface {
	typemap.String[T]
	typemap.Order[T]
}, other T) Matcher[T] {
	return lessThanImpl(tm, "match.LessThanTm", other)
}
