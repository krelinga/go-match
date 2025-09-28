package match

import (
	"fmt"
	"strings"

	"github.com/krelinga/go-match/matchutil"
)

func NewSliceElems[T any](m ...Matcher[T]) SliceElems[T] {
	return SliceElems[T]{M: m}
}

type SliceElems[T any] struct {
	M          []Matcher[T]
	InAnyOrder bool
	Format     func(t T) string
}

func (s SliceElems[T]) Match(got []T) bool {
	if len(got) != len(s.M) {
		return false
	}
	if s.InAnyOrder {
		return s.matchUnordered(got)
	} else {
		return s.matchOrdered(got)
	}
}

func (s SliceElems[T]) matchOrdered(got []T) bool {
	for i, elem := range got {
		if !s.M[i].Match(elem) {
			return false
		}
	}
	return true
}

func (s SliceElems[T]) matchUnordered(got []T) bool {
	used := make([]bool, len(got))
	for _, m := range s.M {
		found := false
		for i, elem := range got {
			if !used[i] && m.Match(elem) {
				used[i] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (s SliceElems[T]) Explain(got []T) string {
	match := s.Match(got)
	var details []string
	lenEq := NewEqual(len(s.M))
	details = append(details, fmt.Sprintf("length:\n%s", matchutil.Indent(Explain(len(got), lenEq), 1)))
	if s.InAnyOrder {
		details = append(details, s.unorderedDetails(got)...)
	} else {
		details = append(details, s.orderedDetails(got)...)
	}
	return matchutil.Explain(match, matchutil.TypeName(s), details...)
}

func (s SliceElems[T]) orderedDetails(got []T) []string {
	var details []string
	for i := range min(len(got), len(s.M)) {
		detail := fmt.Sprintf("index %d:\n%s", i, matchutil.Indent(Explain(got[i], s.M[i]), 1))
		details = append(details, detail)
	}
	return details
}

func (s SliceElems[T]) unorderedDetails(got []T) []string {
	var matched, gotUnmatched []string
	used := make([]bool, len(s.M))
	for gotI, g := range got {
		var found Matcher[T]
		var foundIndex int
		for mI, m := range s.M {
			if !used[mI] && Match(g, m) {
				used[mI] = true
				found = m
				foundIndex = mI
				break
			}
		}
		if found != nil {
			explained := matchutil.Indent(Explain(g, found), 1)
			matched = append(matched, fmt.Sprintf("index %d:\n%s", foundIndex, explained))
		} else {
			value := matchutil.Indent(fmt.Sprintf("%s %s", matchutil.Emoji(false), matchutil.FormatWith(g, s.Format)), 1)
			gotUnmatched = append(gotUnmatched, fmt.Sprintf("index %d:\n%s", gotI, value))
		}
	}
	var mUnmatched []string
	for mI := range s.M {
		if !used[mI] {
			matcherStr := matchutil.Indent(fmt.Sprintf("%s %s", matchutil.Emoji(false), matchutil.Describe(s.M[mI])), 1)
			mUnmatched = append(mUnmatched, fmt.Sprintf("index %d:\n%s", mI, matcherStr))
		}
	}
	var details []string
	if len(matched) > 0 {
		matchedStr := matchutil.Indent(strings.Join(matched, "\n"), 1)
		details = append(details, fmt.Sprintf("%s matched elements:\n%s", matchutil.Emoji(true), matchedStr))
	}
	if len(gotUnmatched) > 0 {
		gotUnmatchedStr := matchutil.Indent(strings.Join(gotUnmatched, "\n"), 1)
		details = append(details, fmt.Sprintf("%s got unmatched elements:\n%s", matchutil.Emoji(false), gotUnmatchedStr))
	}
	if len(mUnmatched) > 0 {
		mUnmatchedStr := matchutil.Indent(strings.Join(mUnmatched, "\n"), 1)
		details = append(details, fmt.Sprintf("%s unmatched matchers:\n%s", matchutil.Emoji(false), mUnmatchedStr))
	}
	return details
}

func NewSliceLen[T any](m Matcher[int]) SliceLen[T] {
	return SliceLen[T]{M: m}
}

type SliceLen[T any] struct {
	M Matcher[int]
}

func (s SliceLen[T]) Match(got []T) bool {
	return s.M.Match(len(got))
}

func (s SliceLen[T]) Explain(got []T) string {
	match := s.Match(got)
	var details []string
	if match {
		details = append(details, "matched length")
	} else {
		details = append(details, "did not match length")
	}
	details = append(details, Explain(len(got), s.M))
	return matchutil.Explain(match, matchutil.TypeName(s), details...)
}

func NewSliceNil[T any]() SliceNil[T] {
	return SliceNil[T]{}
}

type SliceNil[T any] struct {
	Format func(t []T) string
}

func (_ SliceNil[T]) Match(got []T) bool {
	return got == nil
}

func (sn SliceNil[T]) Explain(got []T) string {
	match := sn.Match(got)
	var details []string
	expected := "got == nil"
	if match {
		details = append(details, expected)
	} else {
		actual := fmt.Sprintf("got == %s", matchutil.FormatWith(got, sn.Format))
		details = append(details, matchutil.ActualVsExpected(actual, expected))
	}
	return matchutil.Explain(match, matchutil.TypeName(sn), details...)
}

func NewSliceHas[T any](m Matcher[T]) SliceHas[T] {
	return SliceHas[T]{M: m}
}

type SliceHas[T any] struct {
	M Matcher[T]
}

func (s SliceHas[T]) Match(got []T) bool {
	for _, elem := range got {
		if s.M.Match(elem) {
			return true
		}
	}
	return false
}

func (s SliceHas[T]) Explain(got []T) string {
	match := s.Match(got)
	var details []string
	if match {
		details = append(details, "found matching element")
	} else {
		details = append(details, "no matching element found")
	}
	for i, elem := range got {
		detail := fmt.Sprintf("index %d:\n%s", i, matchutil.Indent(Explain(elem, s.M), 1))
		details = append(details, detail)
	}
	return matchutil.Explain(match, matchutil.TypeName(s), details...)
}