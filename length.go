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

func StringLikeLength[T ~string](matcher Matcher[int]) Matcher[T] {
	tm := typemap.ForStringLike[T]{
		StringFunc: DefaultString[T],
	}
	return lengthImpl(tm, "match.StringLikeLength", matcher)
}

func StringLength(matcher Matcher[int]) Matcher[string] {
	tm := typemap.ForString{
		StringFunc: DefaultString[string],
	}
	return lengthImpl(tm, "match.StringLength", matcher)
}

func SliceLikeLength[T ~[]E, E any](matcher Matcher[int]) Matcher[T] {
	tm := typemap.ForSliceLike[T, E]{
		StringFunc: DefaultString[T],
	}
	return lengthImpl(tm, "match.SliceLikeLength", matcher)
}

func SliceLength[E any](matcher Matcher[int]) Matcher[[]E] {
	tm := typemap.ForSlice[E]{
		StringFunc: DefaultString[[]E],
	}
	return lengthImpl(tm, "match.SliceLength", matcher)
}

func MapLikeLength[T ~map[K]V, K comparable, V any](matcher Matcher[int]) Matcher[T] {
	tm := typemap.ForMapLike[T, K, V]{
		StringFunc: DefaultString[T],
	}
	return lengthImpl(tm, "match.MapLikeLength", matcher)
}

func MapLength[K comparable, V any](matcher Matcher[int]) Matcher[map[K]V] {
	tm := typemap.ForMap[K, V]{
		StringFunc: DefaultString[map[K]V],
	}
	return lengthImpl(tm, "match.MapLength", matcher)
}
