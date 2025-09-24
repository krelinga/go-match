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

type MatchFunc[T any] func(got T) *ResultImpl

type ResultImpl struct {
	Match   bool
	Message string
	Name    string

	Children []*ChildResult
}

func (r *ResultImpl) Matched() bool {
	return r.Match
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
	Result Result
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
	return r.Matched()
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
			return newLeafResult(got == want, got, want, "==", lm.format)
		},
	}
}

func NotEqual[T comparable](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			return newLeafResult(got != want, got, want, "!=", lm.format)
		},
	}
}

func LessThan[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			return newLeafResult(got < want, got, want, "<", lm.format)
		},
	}
}

func GreaterThan[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			return newLeafResult(got > want, got, want, ">", lm.format)
		},
	}
}

func LessThanOrEqual[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			return newLeafResult(got <= want, got, want, "<=", lm.format)
		},
	}
}

func GreaterThanOrEqual[T cmp.Ordered](want T) *LeafMatcher[T] {
	return &LeafMatcher[T]{
		fn: func(lm *LeafMatcher[T], got T) Result {
			return newLeafResult(got >= want, got, want, ">=", lm.format)
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

func newLeafResult[T any](matched bool, got, want T, symbol string, valueFmt func(T) string) leafResult[T] {
	return leafResult[T]{
		matched: matched,
		stringFunc: func(matched bool) string {
			var matchPart string
			if matched {
				matchPart = " "
			} else {
				matchPart = "!"
			}
			return fmt.Sprintf("%s Expected %s %s %s", matchPart, valueFmt(got), symbol, valueFmt(want))
		},
	}
}

type leafResult[T any] struct {
	stringFunc func(matched bool) string
	matched    bool
}

func (lr leafResult[T]) Matched() bool {
	return lr.matched
}

func (lr leafResult[T]) String() string {
	return lr.stringFunc(lr.matched)
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
