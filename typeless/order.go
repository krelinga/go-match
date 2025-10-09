package typeless

import (
	"cmp"
	"fmt"

	"github.com/krelinga/go-match/matchfmt"
)

type OrdFunc func(a, b any) (int, error)

func DefaultOrd[T cmp.Ordered](a, b any) (int, error) {
	aOrd, okA := a.(T)
	if !okA {
		return 0, Error(ErrType, fmt.Sprintf("for a, expected %T, got %T", aOrd, a))
	}
	bOrd, okB := b.(T)
	if !okB {
		return 0, Error(ErrType, fmt.Sprintf("for b, expected %T, got %T", bOrd, b))
	}
	return cmp.Compare(aOrd, bOrd), nil
}

type LessThanMatcher struct {
	Val any
	Ord OrdFunc
	Fmt FmtFunc
}

func LessThan[T cmp.Ordered](val T) LessThanMatcher {
	return LessThanMatcher{
		Val: val,
		Ord: DefaultOrd[T],
		Fmt: DefaultFmt,
	}
}

func (m LessThanMatcher) Match(got any) (Matched, Explanation, error) {
	compare, err := m.Ord(got, m.Val)
	if err != nil {
		return false, "", err
	}
	lessThan := compare < 0
	gotStr, err := m.Fmt(got)
	if err != nil {
		return false, "", err
	}
	valStr, err := m.Fmt(m.Val)
	if err != nil {
		return false, "", err
	}
	var detail string
	expected := fmt.Sprintf("got < %s", valStr)
	if lessThan {
		detail = expected
	} else {
		actual := fmt.Sprintf("got == %s", gotStr)
		detail = matchfmt.ActualVsExpected(actual, expected)
	}
	explanation := matchfmt.Explain(lessThan, "match.LessThanMatcher", detail)
	return Matched(lessThan), Explanation(explanation), nil
}