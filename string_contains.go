package match

import (
	"fmt"
	"strings"

	"github.com/krelinga/go-match/matchfmt"
	"github.com/krelinga/go-typemap"
)

func stringContainsImpl[T ~string](tm typemap.String[T], name string, substr string) Matcher[T] {
	return MatcherFunc[T](func(got T) (match bool, explanation string) {
		strGot := string(got)
		match = strings.Contains(strGot, substr)
		expected := fmt.Sprintf("string contains %q", substr)
		var detail string
		if match {
			detail = expected
		} else {
			actual := fmt.Sprintf("string %s does not contain %q", tm.String(got), substr)
			detail = matchfmt.ActualVsExpected(actual, expected)
		}
		explanation = matchfmt.Explain(match, name, detail)
		return
	})
}

func StringContainsTm[T ~string](tm typemap.String[T], substr string) Matcher[T] {
	return stringContainsImpl(tm, "match.StringContainsTm", substr)
}

func StringLikeContains[T ~string](substr string) Matcher[T] {
	tm := typemap.ForStringLike[T]{
		StringFunc: DefaultString[T],
	}
	return stringContainsImpl(tm, "match.StringLikeContains", substr)
}

func StringContains(substr string) Matcher[string] {
	tm := typemap.ForString{
		StringFunc: DefaultString[string],
	}
	return stringContainsImpl(tm, "match.StringContains", substr)
}