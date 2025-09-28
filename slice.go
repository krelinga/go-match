package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
)

func NewSliceElems[T any](m ...Matcher[T]) SliceElems[T] {
	return SliceElems[T]{M: m}
}

type SliceElems[T any] struct {
	M []Matcher[T]
}

func (s SliceElems[T]) Match(got []T) bool {
	if len(got) != len(s.M) {
		return false
	}
	for i, elem := range got {
		if !s.M[i].Match(elem) {
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
	for i := range min(len(got), len(s.M)) {
		detail := fmt.Sprintf("index %d:\n%s", i, matchutil.Indent(Explain(got[i], s.M[i]), 1))
		details = append(details, detail)
	}
	return matchutil.Explain(match, matchutil.TypeName(s), details...)
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