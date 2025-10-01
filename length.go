package match

import (
	"github.com/krelinga/go-match/matchutil"
	"github.com/krelinga/go-typemap"
)

func lengthImpl[T any](tm typemap.Length[T], name string, matcher Matcher[int]) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		length := tm.Length(got)
		matched, e := matcher.Match(length)
		explanation = matchutil.Explain(matched, name, e)
		return
	})
}

func LengthTm[T any](tm typemap.Length[T], matcher Matcher[int]) Matcher[T] {
	return lengthImpl(tm, "match.LengthTm", matcher)
}

func StringLength[T ~string](matcher Matcher[int]) Matcher[T] {
	return lengthImpl(typemap.ForStringLike[T]{}, "match.StringLength", matcher)
}

func SliceLength[T ~[]E, E any](matcher Matcher[int]) Matcher[T] {
	return lengthImpl(typemap.ForSliceLike[T, E]{}, "match.SliceLength", matcher)
}

func MapLength[T ~map[K]V, K comparable, V any](matcher Matcher[int]) Matcher[T] {
	return lengthImpl(typemap.ForMapLike[T, K, V]{}, "match.MapLength", matcher)
}