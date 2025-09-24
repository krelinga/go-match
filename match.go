package match

import (
	"cmp"
	"fmt"
	"strings"
)

type Matcher[T any] interface {
	Match(got T) *ResultImpl
}

func Match[T any](got T, matcher Matcher[T]) *ResultImpl {
	return matcher.Match(got)
}

type MatchFunc[T any] func(got T) *ResultImpl

type ResultImpl struct {
	Match   bool
	Message string
	Name    string

	Children []*ChildResult
}

func (r *ResultImpl) String() string {
	var matchPart string
	if r.Match {
		matchPart = " "
	} else {
		matchPart = "!"
	}
	summary := fmt.Sprintf("%s %s: %s", matchPart, r.Name, r.Message)
	var sb strings.Builder
	sb.WriteString(summary)
	prefix := fmt.Sprintf("\n%s   ", matchPart)
	for i, child := range r.Children {
		if i == 0 {
			sb.WriteString(fmt.Sprintf("\n%s Children:", matchPart))
		}
		sb.WriteString(prefix)
		sb.WriteString(strings.ReplaceAll(child.String(), "\n", prefix))
	}
	return sb.String()
}

type ChildResult struct {
	Name   string
	Result *ResultImpl
}

func (cr *ChildResult) String() string {
	resultString := cr.Result.String()
	indentedResultString := strings.ReplaceAll(resultString, "\n", "\n  ")
	return fmt.Sprintf("- %s:\n  %s", cr.Name, indentedResultString)
}

func MatchChild[T any](parent *ResultImpl, name string, got T, matcher Matcher[T]) bool {
	r := matcher.Match(got)
	parent.Children = append(parent.Children, &ChildResult{
		Name:   name,
		Result: r,
	})
	return r.Match
}

type matcher[T any] MatchFunc[T]

func (m matcher[T]) Match(got T) *ResultImpl {
	return m(got)
}

func NewMatcher[T any](matchFunc MatchFunc[T]) Matcher[T] {
	return matcher[T](matchFunc)
}

func defaultFormatFunc[T any](v T) string {
	return fmt.Sprintf("%v", v)
}

func Equal[T comparable](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) *ResultImpl {
			return &ResultImpl{
				Name:    "Equal",
				Match:   got == want,
				Message: fmt.Sprintf("Expected %s == %s", lm.format(got), lm.format(want)),
			}
		},
	}
}

func NotEqual[T comparable](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) *ResultImpl {
			return &ResultImpl{
				Name:    "NotEqual",
				Match:   got != want,
				Message: fmt.Sprintf("Expected %s != %s", lm.format(got), lm.format(want)),
			}
		},
	}
}

func LessThan[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) *ResultImpl {
			return &ResultImpl{
				Name:    "LessThan",
				Match:   got < want,
				Message: fmt.Sprintf("Expected %s < %s", lm.format(got), lm.format(want)),
			}
		},
	}
}

func GreaterThan[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) *ResultImpl {
			return &ResultImpl{
				Name:    "GreaterThan",
				Match:   got > want,
				Message: fmt.Sprintf("Expected %s > %s", lm.format(got), lm.format(want)),
			}
		},
	}
}

func LessThanOrEqual[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) *ResultImpl {
			return &ResultImpl{
				Name:    "LessThanOrEqual",
				Match:   got <= want,
				Message: fmt.Sprintf("Expected %s <= %s", lm.format(got), lm.format(want)),
			}
		},
	}
}

func GreaterThanOrEqual[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) *ResultImpl {
			return &ResultImpl{
				Name:    "GreaterThanOrEqual",
				Match:   got >= want,
				Message: fmt.Sprintf("Expected %s >= %s", lm.format(got), lm.format(want)),
			}
		},
	}
}

type LeafMatcher[T comparable] struct {
	fn         func(lm *LeafMatcher[T], got T) *ResultImpl
	formatFunc func(T) string
}

func (lm *LeafMatcher[T]) WithFormatFunc(formatFunc func(T) string) *LeafMatcher[T] {
	lm.formatFunc = formatFunc
	return lm
}

func (lm *LeafMatcher[T]) format(t T) string {
	if lm.formatFunc != nil {
		return lm.formatFunc(t)
	}
	return defaultFormatFunc(t)
}

func (lm *LeafMatcher[T]) Match(got T) *ResultImpl {
	return lm.fn(lm, got)
}

func AllOf[T any](children ...Matcher[T]) Matcher[T] {
	return NewMatcher(func(got T) *ResultImpl {
		r := &ResultImpl{
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
	return NewMatcher(func(got T) *ResultImpl {
		r := &ResultImpl{
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
	return NewMatcher(func(got T) *ResultImpl {
		r := &ResultImpl{
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
