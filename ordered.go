package match

import (
	"cmp"
	"fmt"

	"github.com/krelinga/go-match/matchutil"
)

func NewLessThan[T cmp.Ordered](x T) LessThan[T] {
	return LessThan[T]{X: x}
}

type LessThan[T cmp.Ordered] struct {
	X      T
	Format func(t T) string
}

func (lt LessThan[T]) Match(got T) bool {
	return got < lt.X
}

func (lt LessThan[T]) Explain(got T) string {
	match := lt.Match(got)
	var details []string
	expected := fmt.Sprintf("got < %s", matchutil.FormatWith(lt.X, lt.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got >= %s", matchutil.FormatWith(lt.X, lt.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(lt), details...)
}

func NewLessThanOrEqual[T cmp.Ordered](x T) LessThanOrEqual[T] {
	return LessThanOrEqual[T]{X: x}
}

type LessThanOrEqual[T cmp.Ordered] struct {
	X      T
	Format func(t T) string
}

func (lte LessThanOrEqual[T]) Match(got T) bool {
	return got <= lte.X
}

func (lte LessThanOrEqual[T]) Explain(got T) string {
	match := lte.Match(got)
	var details []string
	expected := fmt.Sprintf("got <= %s", matchutil.FormatWith(lte.X, lte.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got > %s", matchutil.FormatWith(lte.X, lte.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(lte), details...)
}

func NewGreaterThan[T cmp.Ordered](x T) GreaterThan[T] {
	return GreaterThan[T]{X: x}
}

type GreaterThan[T cmp.Ordered] struct {
	X      T
	Format func(t T) string
}

func (gt GreaterThan[T]) Match(got T) bool {
	return got > gt.X
}

func (gt GreaterThan[T]) Explain(got T) string {
	match := gt.Match(got)
	var details []string
	expected := fmt.Sprintf("got > %s", matchutil.FormatWith(gt.X, gt.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got <= %s", matchutil.FormatWith(gt.X, gt.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(gt), details...)
}

func NewGreaterThanOrEqual[T cmp.Ordered](x T) GreaterThanOrEqual[T] {
	return GreaterThanOrEqual[T]{X: x}
}

type GreaterThanOrEqual[T cmp.Ordered] struct {
	X      T
	Format func(t T) string
}

func (gte GreaterThanOrEqual[T]) Match(got T) bool {
	return got >= gte.X
}

func (gte GreaterThanOrEqual[T]) Explain(got T) string {
	match := gte.Match(got)
	var details []string
	expected := fmt.Sprintf("got >= %s", matchutil.FormatWith(gte.X, gte.Format))
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got < %s", matchutil.FormatWith(gte.X, gte.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(gte), details...)
}