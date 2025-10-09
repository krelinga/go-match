package push2

import (
	"fmt"

	"github.com/krelinga/go-match/matchfmt"
)

type CmpFunc[T any] func(a, b T) bool

func DefaultCmp[T comparable]() CmpFunc[T] {
	return func(a, b T) bool {
		return a == b
	}
}

func EqualResult(equal bool, gotStr, wantStr string) Result {
	expected := fmt.Sprintf("got == %s", wantStr)
	var detail string
	if equal {
		detail = expected
	} else {
		actual := fmt.Sprintf("got == %s", gotStr)
		detail = matchfmt.ActualVsExpected(actual, expected)
	}
	return Result{
		Matched: equal,
		Reason:  matchfmt.Explain(equal, "EqualMatcher", detail),
	}
}

type EqualMatcher[T any] struct {
	Val T
	Cmp CmpFunc[T]
	Fmt FmtFunc[T]
}

func (o EqualMatcher[T]) Match(got T) Result {
	return EqualResult(o.Cmp(got, o.Val), o.Fmt(got), o.Fmt(o.Val))
}

func Equal[T comparable](want T) EqualMatcher[T] {
	return EqualMatcher[T]{Val: want, Cmp: DefaultCmp[T](), Fmt: DefaultFmtFunc[T]()}
}

func NotEqualResult(notEqual bool, gotStr, wantStr string) Result {
	expected := fmt.Sprintf("got != %s", wantStr)
	var detail string
	if notEqual {
		detail = expected
	} else {
		actual := fmt.Sprintf("got == %s", gotStr)
		detail = matchfmt.ActualVsExpected(actual, expected)
	}
	return Result{
		Matched: notEqual,
		Reason:  matchfmt.Explain(notEqual, "NotEqualMatcher", detail),
	}
}

type NotEqualMatcher[T any] struct {
	Val T
	Cmp CmpFunc[T]
	Fmt FmtFunc[T]
}

func (o NotEqualMatcher[T]) Match(got T) Result {
	return NotEqualResult(!o.Cmp(got, o.Val), o.Fmt(got), o.Fmt(o.Val))
}

func NotEqual[T comparable](want T) NotEqualMatcher[T] {
	return NotEqualMatcher[T]{Val: want, Cmp: DefaultCmp[T](), Fmt: DefaultFmtFunc[T]()}
}
