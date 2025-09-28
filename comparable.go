package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
)

func NewEqual[T comparable](x T) Equal[T] {
	return Equal[T]{X: x}
}

type Equal[T comparable] struct {
	X      T
	Format func(t T) string
}

func (e Equal[T]) Match(got T) bool {
	return got == e.X
}

func (e Equal[T]) Explain(got T) string {
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

func (e Equal[T]) String() string {
	return matchutil.StructRepr(
		e,
		matchutil.StructField("X", matchutil.FormatWith(e.X, e.Format)),
		matchutil.StructField("Format", matchutil.Repr(e.Format)),
	)
}

func NewNotEqual[T comparable](x T) NotEqual[T] {
	return NotEqual[T]{X: x}
}

type NotEqual[T comparable] struct {
	X      T
	Format func(t T) string
}

func (ne NotEqual[T]) Match(got T) bool {
	return got != ne.X
}

func (ne NotEqual[T]) Explain(got T) string {
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
