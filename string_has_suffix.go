package match

import (
	"fmt"
	"strings"

	"github.com/krelinga/go-match/matchfmt"
	"github.com/krelinga/go-typemap"
)

func stringLikeHasSuffixImpl[T ~string](tm typemap.String[T], name string, suffix string) Matcher[T] {
	return MatcherFunc[T](func(got T) (match bool, explanation string) {
		strGot := string(got)
		match = strings.HasSuffix(strGot, suffix)
		expected := fmt.Sprintf("string ends with %q", suffix)
		var detail string
		if match {
			detail = expected
		} else {
			actual := fmt.Sprintf("string %s does not end with %q", tm.String(got), suffix)
			detail = matchfmt.ActualVsExpected(actual, expected)
		}
		explanation = matchfmt.Explain(match, name, detail)
		return
	})
}

func StringLikeHasSuffixTm[T ~string](tm typemap.String[T], suffix string) Matcher[T] {
	return stringLikeHasSuffixImpl(tm, "match.StringLikeHasSuffixTm", suffix)
}

func StringLikeHasSuffix[T ~string](suffix string) Matcher[T] {
	tm := typemap.ForStringLike[T]{
		StringFunc: DefaultString[T],
	}
	return stringLikeHasSuffixImpl(tm, "match.StringLikeHasSuffix", suffix)
}

func StringHasSuffix(suffix string) Matcher[string] {
	tm := typemap.ForString{
		StringFunc: DefaultString[string],
	}
	return stringLikeHasSuffixImpl(tm, "match.StringHasSuffix", suffix)
}
