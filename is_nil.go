package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
	"github.com/krelinga/go-typemap"
)

func isNilImpl[T any](tm interface {
	typemap.IsNil[T]
	typemap.String[T]
}, name string) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = tm.IsNil(got)
		expected := "got == nil"
		if matched {
			explanation = expected
		} else {
			actual := fmt.Sprintf("got = %s", tm.String(got))
			explanation = matchutil.Explain(matched, name, matchutil.ActualVsExpected(actual, expected))
		}
		return
	})
}

func IsNilTm[T any](tm interface {
	typemap.IsNil[T]
	typemap.String[T]
}) Matcher[T] {
	return isNilImpl(tm, "match.IsNilTm")
}

func SliceIsNil[T ~[]E, E any]() Matcher[T] {
	tm := typemap.ForSliceLike[T, E]{
		StringFunc: DefaultString[T](),
	}
	return isNilImpl(tm, "match.SliceIsNil")
}

func MapIsNil[T ~map[K]V, K comparable, V any]() Matcher[T] {
	tm := typemap.ForMapLike[T, K, V]{
		StringFunc: DefaultString[T](),
	}
	return isNilImpl(tm, "match.MapIsNil")
}

func PointerIsNil[T any]() Matcher[*T] {
	tm := typemap.ForPointer[T]{
		StringFunc: DefaultString[*T](),
	}
	return isNilImpl(tm, "match.PointerIsNil")
}
