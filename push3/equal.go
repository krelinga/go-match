package push3

import (
	"fmt"

	"github.com/krelinga/go-match/matchfmt"
)

type Cmp[T any] func(a, b T) bool

func DefaultCmp[T comparable]() Cmp[T] {
	return func(a, b T) bool {
		return a == b
	}
}

type Equal[T any] struct {
	Val T
	Cmp Cmp[T]
	Fmt Fmt[T]
}

func (o Equal[T]) M(got T) Result {
	equal := o.Cmp(got, o.Val)
	var detail string
	expected := fmt.Sprintf("got == %s", o.Fmt(o.Val))
	if equal {
		detail = expected
	} else {
		actual := fmt.Sprintf("got == %s", o.Fmt(got))
		detail = matchfmt.ActualVsExpected(actual, expected)
	}
	return Result{
		Matched:     equal,
		Explanation: matchfmt.Explain(equal, "Equal", detail),
	}
}

func EqualM[T comparable](want T) M[T] {
	return Equal[T]{Val: want, Cmp: DefaultCmp[T](), Fmt: DefaultFmt[T]()}.M
}

type NotEqual[T any] struct {
	Val T
	Cmp Cmp[T]
	Fmt Fmt[T]
}

func (o NotEqual[T]) M(got T) Result {
	notEqual := !o.Cmp(got, o.Val)
	var detail string
	expected := fmt.Sprintf("got != %s", o.Fmt(o.Val))
	if notEqual {
		detail = expected
	} else {
		actual := fmt.Sprintf("got == %s", o.Fmt(got))
		detail = matchfmt.ActualVsExpected(actual, expected)
	}
	return Result{
		Matched:     notEqual,
		Explanation: matchfmt.Explain(notEqual, "NotEqual", detail),
	}
}

func NotEqualM[T comparable](want T) M[T] {
	return NotEqual[T]{Val: want, Cmp: DefaultCmp[T](), Fmt: DefaultFmt[T]()}.M
}
