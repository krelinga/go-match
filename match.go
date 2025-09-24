package match

import (
	"fmt"
	"strings"
)

type Matcher[T any] interface {
	Match(got T) *Result
}

type LeafMatcher[T any] interface {
	Matcher[T]
	WithFormatFunc(FormatFunc[T]) LeafMatcher[T]
}

func Match[T any](got T, matcher Matcher[T]) *Result {
	return matcher.Match(got)
}

type FormatFunc[T any] func(T) string

type MatchFunc[T any] func(got T) *Result

type LeafMatchFunc[T any] func(got T, formatFunc FormatFunc[T]) *Result

type Result struct {
	Match   bool
	Message string
	Name    string

	Children []*ChildResult
}

func (r *Result) String() string {
	var matchPart string
	if r.Match {
		matchPart = " "
	} else {
		matchPart = "!"
	}
	summary := fmt.Sprintf("%s %s: %s", matchPart, r.Name, r.Message)
	var sb strings.Builder
	sb.WriteString(summary)
	for i, child := range r.Children {
		if i == 0 {
			sb.WriteString("\n  Children:")
		}
		sb.WriteString("\n    ")
		sb.WriteString(strings.ReplaceAll(child.String(), "\n", "\n    "))
	}
	return sb.String()
}

type ChildResult struct {
	Name   string
	Result *Result
}

func (cr *ChildResult) String() string {
	resultString := cr.Result.String()
	indentedResultString := strings.ReplaceAll(resultString, "\n", "\n  ")
	return fmt.Sprintf("- %s:\n  %s", cr.Name, indentedResultString)
}

func MatchChild[T any](parent *Result, name string, got T, matcher Matcher[T]) bool {
	r := matcher.Match(got)
	parent.Children = append(parent.Children, &ChildResult{
		Name:   name,
		Result: r,
	})
	return r.Match
}

type matcher[T any] MatchFunc[T]

func (m matcher[T]) Match(got T) *Result {
	return m(got)
}

func NewMatcher[T any](matchFunc MatchFunc[T]) Matcher[T] {
	return matcher[T](matchFunc)
}

type leafMatcher[T any] struct {
	leafMatchFunc LeafMatchFunc[T]
	formatFunc    FormatFunc[T]
}

func defaultFormatFunc[T any](v T) string {
	return fmt.Sprintf("%v", v)
}

func (lm *leafMatcher[T]) Match(got T) *Result {
	var formatFunc FormatFunc[T]
	if lm.formatFunc != nil {
		formatFunc = lm.formatFunc
	} else {
		formatFunc = defaultFormatFunc[T]
	}
	return lm.leafMatchFunc(got, formatFunc)
}

func (lm *leafMatcher[T]) WithFormatFunc(formatFunc FormatFunc[T]) LeafMatcher[T] {
	lm.formatFunc = formatFunc
	return lm
}

func NewLeafMatcher[T any](leafMatchFunc LeafMatchFunc[T]) *leafMatcher[T] {
	return &leafMatcher[T]{
		leafMatchFunc: leafMatchFunc,
	}
}

func Equals[T comparable](want T) LeafMatcher[T] {
	return NewLeafMatcher(func(got T, formatFunc FormatFunc[T]) *Result {
		match := want == got
		return &Result{
			Name:    "Equals",
			Match:   match,
			Message: fmt.Sprintf("Expected %s == %s", formatFunc(want), formatFunc(got)),
		}
	})
}

func AllOf[T any](children ...Matcher[T]) Matcher[T] {
	return NewMatcher(func(got T) *Result {
		r := &Result{
			Name:  "AllOf",
			Match: true,
		}
		unmatched := []string{}
		for i, child := range children {
			childStr := fmt.Sprintf("%d", i)
			if !MatchChild(r, childStr, got, child) {
				r.Match = false
				unmatched = append(unmatched, childStr)
			}
		}
		if r.Match {
			r.Message = "all matched"
		} else {
			r.Message = fmt.Sprintf("unmatched indices: %s", strings.Join(unmatched, ", "))
		}
		return r
	})
}

func AnyOf[T any](children ...Matcher[T]) Matcher[T] {
	return NewMatcher(func(got T) *Result {
		r := &Result{
			Name:  "AnyOf",
			Match: false,
		}
		matched := []string{}
		for i, child := range children {
			childStr := fmt.Sprintf("%d", i)
			if MatchChild(r, childStr, got, child) {
				r.Match = true
				matched = append(matched, childStr)
			}
		}
		if r.Match {
			r.Message = fmt.Sprintf("matched indices: %s", strings.Join(matched, ", "))
		} else {
			r.Message = "no matches"
		}
		return r
	})
}

func Not[T any](child Matcher[T]) Matcher[T] {
	return NewMatcher(func(got T) *Result {
		r := &Result{
			Name: "Not",
		}
		r.Match = !MatchChild(r, "child", got, child)
		if r.Match {
			r.Message = "child did not match"
		} else {
			r.Message = "child matched"
		}
		return r
	})
}
