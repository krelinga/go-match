package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
)

type Matcher[T any] interface {
	Match(got T) bool
}

type Explainer[T any] interface {
	Explain(got T) string
}

func Match[T any](got T, matcher Matcher[T]) bool {
	return matcher.Match(got)
}

func Explain[T any](got T, matcher Matcher[T]) string {
	if explainer, ok := matcher.(Explainer[T]); ok {
		return explainer.Explain(got)
	}
	return matchutil.Explain(matcher.Match(got), matchutil.TypeName(matcher))
}

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

func NewAllOf[T any](m ...Matcher[T]) AllOf[T] {
	return AllOf[T]{M: m}
}

type AllOf[T any] struct {
	M []Matcher[T]
}

func (a AllOf[T]) Match(got T) bool {
	for _, m := range a.M {
		if !m.Match(got) {
			return false
		}
	}
	return true
}

func (a AllOf[T]) Explain(got T) string {
	match := a.Match(got)
	var details []string
	if match {
		details = append(details, "matched all conditions")
	}
	if !match {
		details = append(details, "did not match all conditions")
	}
	for i, m := range a.M {
		detail := fmt.Sprintf("index %d:\n%s", i, matchutil.Indent(Explain(got, m), 1))
		details = append(details, detail)
	}
	return matchutil.Explain(match, matchutil.TypeName(a), details...)
}

type AnyOf[T any] struct {
	M []Matcher[T]
}

func NewAnyOf[T any](m ...Matcher[T]) AnyOf[T] {
	return AnyOf[T]{M: m}
}

func (a AnyOf[T]) Match(got T) bool {
	for _, m := range a.M {
		if m.Match(got) {
			return true
		}
	}
	return false
}

func (a AnyOf[T]) Explain(got T) string {
	match := a.Match(got)
	var details []string
	if match {
		details = append(details, "matched at least one condition")
	} else {
		details = append(details, "did not match any condition")
	}
	for i, m := range a.M {
		detail := fmt.Sprintf("index %d:\n%s", i, matchutil.Indent(Explain(got, m), 1))
		details = append(details, detail)
	}
	return matchutil.Explain(match, matchutil.TypeName(a), details...)
}

func NewWhenDeref[T any](m Matcher[T]) WhenDeref[T] {
	return WhenDeref[T]{M: m}
}

type WhenDeref[T any] struct {
	M Matcher[T]
}

func (p WhenDeref[T]) Match(got *T) bool {
	if got == nil {
		return false
	}
	return p.M.Match(*got)
}

func (p WhenDeref[T]) Explain(got *T) string {
	match := p.Match(got)
	var details []string
	if got == nil {
		details = append(details, "got == nil")
	} else {
		details = append(details, "got != nil", Explain(*got, p.M))
	}
	return matchutil.Explain(match, matchutil.TypeName(p), details...)
}
