package typeless

import (
	"fmt"
	"reflect"

	"github.com/krelinga/go-match/matchfmt"
)

type lenMatcher struct {
	matcher Matcher
}

func Len(matcher Matcher) Matcher {
	return lenMatcher{matcher: matcher}
}

func (m lenMatcher) Match(got any) (Matched, Explanation, error) {
	var length int
	val := reflect.ValueOf(got)
	if !val.IsValid() {
		return false, "", Error(ErrValue, "got value is invalid")
	}
	switch val.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map, reflect.Chan:
		length = val.Len()
	default:
		return false, "", Error(ErrType, fmt.Sprintf("Len matcher requires array, slice, string, or map.  Got type %q, which is kind %s", val.Type(), val.Kind()))
	}
	matched, explanation, err := m.matcher.Match(length)
	if err != nil {
		return false, "", err
	}
	return matched, Explanation(matchfmt.Explain(bool(matched), "match.Len", string(explanation))), nil
}
