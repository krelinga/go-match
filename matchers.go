package match

import (
	"cmp"
	"fmt"
)

func Equals[T comparable](expected T) Matcher[T] {
	return &equalsMatcher[T]{expected: expected}
}

type equalsMatcher[T comparable] struct {
	expected T
}

func (m *equalsMatcher[T]) Name() string {
	return "Equals"
}

func (m *equalsMatcher[T]) Match(input T, reporter Reporter) {
	if input != m.expected {
		reporter.Report(fmt.Sprintf("%v != %v", input, m.expected))
	}
}

func LessThan[T cmp.Ordered](limit T) Matcher[T] {
	return &lessThanMatcher[T]{limit: limit}
}

type lessThanMatcher[T cmp.Ordered] struct {
	limit T
}

func (m *lessThanMatcher[T]) Name() string {
	return "LessThan"
}

func (m *lessThanMatcher[T]) Match(input T, reporter Reporter) {
	if input >= m.limit {
		reporter.Report(fmt.Sprintf("%v >= %v", input, m.limit))
	}
}

func GreaterThan[T cmp.Ordered](limit T) Matcher[T] {
	return &greaterThanMatcher[T]{limit: limit}
}

type greaterThanMatcher[T cmp.Ordered] struct {
	limit T
}

func (m *greaterThanMatcher[T]) Name() string {
	return "GreaterThan"
}

func (m *greaterThanMatcher[T]) Match(input T, reporter Reporter) {
	if input <= m.limit {
		reporter.Report(fmt.Sprintf("%v <= %v", input, m.limit))
	}
}

func AllOf[T any](matchers ...Matcher[T]) Matcher[T] {
	return &allOfMatcher[T]{matchers: matchers}
}

type allOfMatcher[T any] struct {
	matchers []Matcher[T]
}

func (m *allOfMatcher[T]) Name() string {
	return "AllOf"
}

func (m *allOfMatcher[T]) Match(input T, reporter Reporter) {
	for _, matcher := range m.matchers {
		matcher.Match(input, reporter.Child(matcher))
	}
}
