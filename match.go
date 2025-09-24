package match

import (
	"fmt"
	"strings"
)

type MatchFunc[T any] func(want, got T) *Result

type Result struct {
	Match bool
	Message string
	Name string

	Children []*ChildResult
}

type ChildResult struct {
	Name string
	Result *Result
}

func MatchChild[T any](parent *Result, name string, want, got T, fn MatchFunc[T]) bool {
	r := fn(want, got)
	parent.Children = append(parent.Children, &ChildResult{
		Name: name,
		Result: r,
	})
	return r.Match
}

func Equals[T comparable](want, got T) *Result {
	return &Result{
		Name: "Equals",
		Match: want == got,
		Message: fmt.Sprintf("want: %v, got: %v", want, got),
	}
}

func AllOf[T any](children ...MatchFunc[T]) MatchFunc[T] {
	return func(want, got T) *Result {
		r := &Result{
			Name: "AllOf",
			Match: true,
		}
		unmatched := []string{}
		for i, child := range children {
			childStr := fmt.Sprintf("%d", i)
			if !MatchChild(r, childStr, want, got, child) {
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
	}
}

func AnyOf[T any](children ...MatchFunc[T]) MatchFunc[T] {
	return func(want, got T) *Result {
		r := &Result{
			Name: "AnyOf",
			Match: false,
		}
		matched := []string{}
		for i, child := range children {
			childStr := fmt.Sprintf("%d", i)
			if MatchChild(r, childStr, want, got, child) {
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
	}
}

func Not[T any](child MatchFunc[T]) MatchFunc[T] {
	return func(want, got T) *Result {
		r := &Result{
			Name: "Not",
		}
		r.Match = !MatchChild(r, "child", want, got, child)
		if r.Match {
			r.Message = "child did not match"
		} else {
			r.Message = "child matched"
		}
		return r
	}
}