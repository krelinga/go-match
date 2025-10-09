package typeless

import (
	"fmt"
	"reflect"

	"github.com/krelinga/go-match/matchfmt"
)

type emptyMatcher struct{}

func Empty() Matcher {
	return emptyMatcher{}
}

func (m emptyMatcher) Match(got any) (Matched, Explanation, error) {
	val := reflect.ValueOf(got)
	if !val.IsValid() {
		return false, "", Error(ErrValue, "got value is invalid")
	}
	switch val.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map, reflect.Chan:
		length := val.Len()
		matched := Matched(length == 0)
		var detail string
		expected := "got is empty"
		if matched {
			detail = expected
		} else {
			actual := fmt.Sprintf("got has length %d", length)
			detail = matchfmt.ActualVsExpected(actual, expected)
		}
		explanation := Explanation(matchfmt.Explain(bool(matched), "match.Empty", detail))
		return matched, explanation, nil
	default:
		return false, "", Error(ErrType, "Empty matcher requires array, slice, string, map, or chan")
	}
}