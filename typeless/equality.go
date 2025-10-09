package typeless

import (
	"fmt"

	"github.com/krelinga/go-match/matchfmt"
)

// Returns true if A and B are considered equal.
type CompFunc func(a, b any) (bool, error)

func DefaultComp[T comparable](a, b any) (bool, error) {
	aComp, okA := a.(T)
	if !okA {
		return false, Error(ErrType, fmt.Sprintf("for a, expected %T, got %T", aComp, a))
	}
	bComp, okB := b.(T)
	if !okB {
		return false, Error(ErrType, fmt.Sprintf("for b, expected %T, got %T", bComp, b))
	}
	return aComp == bComp, nil
}

type EqualMatcher struct {
	Val  any
	Comp CompFunc
	Fmt  FmtFunc
}

func Equal[T comparable](val T) EqualMatcher {
	return EqualMatcher{
		Val:  val,
		Comp: DefaultComp[T],
		Fmt:  DefaultFmt,
	}
}

func (m EqualMatcher) Match(got any) (Matched, Explanation, error) {
	equal, err := m.Comp(got, m.Val)
	if err != nil {
		return false, "", err
	}
	gotStr, err := m.Fmt(got)
	if err != nil {
		return false, "", err
	}
	valStr, err := m.Fmt(m.Val)
	if err != nil {
		return false, "", err
	}
	var detail string
	expected := fmt.Sprintf("got == %s", valStr)
	if equal {
		detail = expected
	} else {
		actual := fmt.Sprintf("got == %s", gotStr)
		detail = matchfmt.ActualVsExpected(actual, expected)
	}
	explanation := matchfmt.Explain(equal, "match.EqualMatcher", detail)
	return Matched(equal), Explanation(explanation), nil
}

type NotEqualMatcher struct {
	Val  any
	Comp CompFunc
	Fmt  FmtFunc
}

func NotEqual[T comparable](val T) NotEqualMatcher {
	return NotEqualMatcher{
		Val:  val,
		Comp: DefaultComp[T],
		Fmt:  DefaultFmt,
	}
}

func (m NotEqualMatcher) Match(got any) (Matched, Explanation, error) {
	equal, err := m.Comp(got, m.Val)
	if err != nil {
		return false, "", err
	}
	gotStr, err := m.Fmt(got)
	if err != nil {
		return false, "", err
	}
	valStr, err := m.Fmt(m.Val)
	if err != nil {
		return false, "", err
	}
	var detail string
	expected := fmt.Sprintf("got != %s", valStr)
	if !equal {
		detail = expected
	} else {
		actual := fmt.Sprintf("got == %s", gotStr)
		detail = matchfmt.ActualVsExpected(actual, expected)
	}
	explanation := matchfmt.Explain(!equal, "match.NotEqualMatcher", detail)
	return Matched(!equal), Explanation(explanation), nil
}
