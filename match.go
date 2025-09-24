package match

import (
	"cmp"
	"fmt"
	"strings"
)

type Matcher[T any] interface {
	Match(got T) Result
}

type Result interface {
	Matched() bool
	String() string
}

func Match[T any](got T, matcher Matcher[T]) Result {
	return matcher.Match(got)
}

type MatchFunc[T any] func(got T) Result

func matchChar(matched bool) string {
	if matched {
		return " "
	}
	return "!"
}

type SimpleResult struct {
	message string
	matched bool
}

func NewSimpleResult(matched bool, message string) SimpleResult {
	return SimpleResult{
		message: message,
		matched: matched,
	}
}

func (sr SimpleResult) Matched() bool {
	return sr.matched
}

func (sr SimpleResult) String() string {
	prefix := matchChar(sr.matched)
	return fmt.Sprintf("%s %s", prefix, sr.message)
}

type matcher[T any] MatchFunc[T]

func (m matcher[T]) Match(got T) Result {
	return m(got)
}

func NewMatcher[T any](matchFunc MatchFunc[T]) Matcher[T] {
	return matcher[T](matchFunc)
}

func Equal[T comparable](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			matched := got == want
			message := fmt.Sprintf("Expected %s == %s", lm.format(got), lm.format(want))
			return NewSimpleResult(matched, message)
		},
	}
}

func NotEqual[T comparable](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			matched := got != want
			message := fmt.Sprintf("Expected %s != %s", lm.format(got), lm.format(want))
			return NewSimpleResult(matched, message)
		},
	}
}

func LessThan[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			matched := got < want
			message := fmt.Sprintf("Expected %s < %s", lm.format(got), lm.format(want))
			return NewSimpleResult(matched, message)
		},
	}
}

func GreaterThan[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			matched := got > want
			message := fmt.Sprintf("Expected %s > %s", lm.format(got), lm.format(want))
			return NewSimpleResult(matched, message)
		},
	}
}

func LessThanOrEqual[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			matched := got <= want
			message := fmt.Sprintf("Expected %s <= %s", lm.format(got), lm.format(want))
			return NewSimpleResult(matched, message)
		},
	}
}

func GreaterThanOrEqual[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			matched := got >= want
			message := fmt.Sprintf("Expected %s >= %s", lm.format(got), lm.format(want))
			return NewSimpleResult(matched, message)
		},
	}
}

type LeafMatcher[T any] struct {
	fn         func(*LeafMatcher[T], T) Result
	formatFunc func(T) string
}

func (lm *LeafMatcher[T]) WithFormatFunc(formatFunc func(T) string) *LeafMatcher[T] {
	lm.formatFunc = formatFunc
	return lm
}

func (lm *LeafMatcher[T]) Match(got T) Result {
	return lm.fn(lm, got)
}

func (lm *LeafMatcher[T]) format(t T) string {
	if lm.formatFunc != nil {
		return lm.formatFunc(t)
	}
	return fmt.Sprintf("%v", t)
}

func AllOf[T any](children ...Matcher[T]) Matcher[T] {
	return NewMatcher(func(got T) Result {
		childResults := make([]Result, len(children))
		matched := true
		for i, child := range children {
			childResults[i] = child.Match(got)
			matched = matched && childResults[i].Matched()
		}
		summary := "Expected all of the following to match:"
		return NewParentResult(matched, summary, childResults...)
	})
}

func AnyOf[T any](children ...Matcher[T]) Matcher[T] {
	return NewMatcher(func(got T) Result {
		childResults := make([]Result, len(children))
		matched := false
		for i, child := range children {
			childResults[i] = child.Match(got)
			matched = matched || childResults[i].Matched()
		}
		summary := "Expected any of the following to match:"
		return NewParentResult(matched, summary, childResults...)
	})
}

func Not[T any](child Matcher[T]) Matcher[T] {
	return NewMatcher(func(got T) Result {
		childResult := child.Match(got)
		matched := !childResult.Matched()
		summary := "Expected the following to not match:"
		return NewParentResult(matched, summary, childResult)
	})
}

type ParentResult struct {
	children []Result
	summary  string
	matched  bool
}

func NewParentResult(matched bool, summary string, children ...Result) ParentResult {
	return ParentResult{
		children: children,
		summary:  summary,
		matched:  matched,
	}
}

func (pr ParentResult) Matched() bool {
	return pr.matched
}

func (pr ParentResult) String() string {
	matchedPart := matchChar(pr.matched)
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s %s", matchedPart, pr.summary))
	prefix := fmt.Sprintf("\n%s   ", matchedPart)
	for _, child := range pr.children {
		sb.WriteString(prefix)
		sb.WriteString(strings.ReplaceAll(child.String(), "\n", prefix))
	}
	return sb.String()
}