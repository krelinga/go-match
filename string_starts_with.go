package match

import (
	"fmt"
	"strings"

	"github.com/krelinga/go-match/matchfmt"
	"github.com/krelinga/go-typemap"
)

func stringLikeStartsWithImpl[T ~string](tm typemap.String[T], name string, prefix string) Matcher[T] {
	return MatcherFunc[T](func(got T) (match bool, explanation string) {
		strGot := string(got)
		match = strings.HasPrefix(strGot, prefix)
		expected := fmt.Sprintf("string starts with %q", prefix)
		var detail string
		if match {
			detail = expected
		} else {
			actual := fmt.Sprintf("string %s does not start with %q", tm.String(got), prefix)
			detail = matchfmt.ActualVsExpected(actual, expected)
		}
		explanation = matchfmt.Explain(match, name, detail)
		return
	})
}

func StringLikeStartsWithTm[T ~string](tm typemap.String[T], prefix string) Matcher[T] {
	return stringLikeStartsWithImpl(tm, "match.StringLikeStartsWithTm", prefix)
}

func StringLikeStartsWith[T ~string](prefix string) Matcher[T] {
	tm := typemap.ForStringLike[T]{
		StringFunc: DefaultString[T],
	}
	return stringLikeStartsWithImpl(tm, "match.StringLikeStartsWith", prefix)
}

func StringStartsWith(prefix string) Matcher[string] {
	tm := typemap.ForString{
		StringFunc: DefaultString[string],
	}
	return stringLikeStartsWithImpl(tm, "match.StringStartsWith", prefix)
}