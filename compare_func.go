package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
)

func NewEqualFunc[T any](x T, fn func(x, y T) int) EqualFunc[T] {
	return EqualFunc[T]{X: x, Func: fn}
}

type EqualFunc[T any] struct {
	X      T
	Format func(t T) string
	Func   func(x, y T) int
}

func (e EqualFunc[T]) Match(got T) bool {
	if e.Func == nil {
		panic("EqualFunc.Func is nil")
	}
	return e.Func(got, e.X) == 0
}

func (e EqualFunc[T]) Explain(got T) string {
	match := e.Match(got)
	var details []string
	expected := fmt.Sprintf("got == %s", matchutil.FormatWith(e.X, e.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got == %s", matchutil.FormatWith(got, e.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(e), details...)
}

func NewNotEqualFunc[T any](x T, fn func(x, y T) int) NotEqualFunc[T] {
	return NotEqualFunc[T]{X: x, Func: fn}
}

type NotEqualFunc[T any] struct {
	X      T
	Format func(t T) string
	Func   func(x, y T) int
}

func (ne NotEqualFunc[T]) Match(got T) bool {
	if ne.Func == nil {
		panic("NotEqualFunc.Func is nil")
	}
	return ne.Func(got, ne.X) != 0
}

func (ne NotEqualFunc[T]) Explain(got T) string {
	match := ne.Match(got)
	var details []string
	expected := fmt.Sprintf("got != %s", matchutil.FormatWith(ne.X, ne.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got == %s", matchutil.FormatWith(got, ne.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(ne), details...)
}

func NewLessThanFunc[T any](x T, fn func(x, y T) int) LessThanFunc[T] {
	return LessThanFunc[T]{X: x, Func: fn}
}

type LessThanFunc[T any] struct {
	X      T
	Format func(t T) string
	Func   func(x, y T) int
}

func (lt LessThanFunc[T]) Match(got T) bool {
	if lt.Func == nil {
		panic("LessThanFunc.Func is nil")
	}
	return lt.Func(got, lt.X) < 0
}

func (lt LessThanFunc[T]) Explain(got T) string {
	match := lt.Match(got)
	var details []string
	expected := fmt.Sprintf("got < %s", matchutil.FormatWith(lt.X, lt.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got == %s", matchutil.FormatWith(got, lt.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(lt), details...)
}

func NewLessThanOrEqualFunc[T any](x T, fn func(x, y T) int) LessThanOrEqualFunc[T] {
	return LessThanOrEqualFunc[T]{X: x, Func: fn}
}

type LessThanOrEqualFunc[T any] struct {
	X      T
	Format func(t T) string
	Func   func(x, y T) int
}

func (lte LessThanOrEqualFunc[T]) Match(got T) bool {
	if lte.Func == nil {
		panic("LessThanOrEqualFunc.Func is nil")
	}
	return lte.Func(got, lte.X) <= 0
}

func (lte LessThanOrEqualFunc[T]) Explain(got T) string {
	match := lte.Match(got)
	var details []string
	expected := fmt.Sprintf("got <= %s", matchutil.FormatWith(lte.X, lte.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got == %s", matchutil.FormatWith(got, lte.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(lte), details...)
}

func NewGreaterThanFunc[T any](x T, fn func(x, y T) int) GreaterThanFunc[T] {
	return GreaterThanFunc[T]{X: x, Func: fn}
}

type GreaterThanFunc[T any] struct {
	X      T
	Format func(t T) string
	Func   func(x, y T) int
}

func (gt GreaterThanFunc[T]) Match(got T) bool {
	if gt.Func == nil {
		panic("GreaterThanFunc.Func is nil")
	}
	return gt.Func(got, gt.X) > 0
}

func (gt GreaterThanFunc[T]) Explain(got T) string {
	match := gt.Match(got)
	var details []string
	expected := fmt.Sprintf("got > %s", matchutil.FormatWith(gt.X, gt.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got == %s", matchutil.FormatWith(got, gt.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(gt), details...)
}

func NewGreaterThanOrEqualFunc[T any](x T, fn func(x, y T) int) GreaterThanOrEqualFunc[T] {
	return GreaterThanOrEqualFunc[T]{X: x, Func: fn}
}

type GreaterThanOrEqualFunc[T any] struct {
	X      T
	Format func(t T) string
	Func   func(x, y T) int
}

func (gte GreaterThanOrEqualFunc[T]) Match(got T) bool {
	if gte.Func == nil {
		panic("GreaterThanOrEqualFunc.Func is nil")
	}
	return gte.Func(got, gte.X) >= 0
}

func (gte GreaterThanOrEqualFunc[T]) Explain(got T) string {
	match := gte.Match(got)
	var details []string
	expected := fmt.Sprintf("got >= %s", matchutil.FormatWith(gte.X, gte.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got == %s", matchutil.FormatWith(got, gte.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(gte), details...)
}