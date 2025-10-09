package typeless

import (
	"fmt"
	"reflect"

	"github.com/krelinga/go-match/matchfmt"
)

type NilMatcher struct {
	Fmt FmtFunc
}

func Nil() NilMatcher {
	return NilMatcher{
		Fmt: DefaultFmt,
	}
}

func (m NilMatcher) Match(got any) (Matched, Explanation, error) {
	val := reflect.ValueOf(got)
	var matched Matched
	gotStr, err := m.Fmt(got)
	if err != nil {
		return false, "", err
	}
	if !val.IsValid() {
		// Nil interface
		matched = true
	} else {
		switch val.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
			matched = Matched(val.IsNil())
		default:
			return false, "", Error(ErrType, fmt.Sprintf("Nil matcher requires chan, func, interface, map, pointer, or slice.  Got type %q kind %s", val.Type(), val.Kind()))
		}
	}
	expected := "got == nil"
	var detail string
	if matched {
		detail = expected
	} else {
		actual := fmt.Sprintf("got == %s", gotStr)
		detail = matchfmt.ActualVsExpected(actual, expected)
	}
	explanation := Explanation(matchfmt.Explain(bool(matched), "match.Nil", detail))
	return matched, explanation, nil
}