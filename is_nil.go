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
		var detail string
		if matched {
			detail = expected
		} else {
			actual := fmt.Sprintf("got = %s", tm.String(got))
			detail = matchutil.ActualVsExpected(actual, expected)
		}
		explanation = matchutil.Explain(matched, name, detail)
		return
	})
}

func IsNilTm[T any](tm interface {
	typemap.IsNil[T]
	typemap.String[T]
}) Matcher[T] {
	return isNilImpl(tm, "match.IsNilTm")
}

func SliceLikeIsNil[T ~[]E, E any]() Matcher[T] {
	tm := typemap.ForSliceLike[T, E]{
		StringFunc: DefaultString[T](),
	}
	return isNilImpl(tm, "match.SliceLikeIsNil")
}

func SliceIsNil[E any]() Matcher[[]E] {
	tm := typemap.ForSlice[E]{
		StringFunc: DefaultString[[]E](),
	}
	return isNilImpl(tm, "match.SliceIsNil")
}

func MapLikeIsNil[T ~map[K]V, K comparable, V any]() Matcher[T] {
	tm := typemap.ForMapLike[T, K, V]{
		StringFunc: DefaultString[T](),
	}
	return isNilImpl(tm, "match.MapLikeIsNil")
}

func MapIsNil[K comparable, V any]() Matcher[map[K]V] {
	tm := typemap.ForMap[K, V]{
		StringFunc: DefaultString[map[K]V](),
	}
	return isNilImpl(tm, "match.MapIsNil")
}

func PointerIsNil[T any]() Matcher[*T] {
	tm := typemap.ForPointer[T]{
		StringFunc: DefaultString[*T](),
	}
	return isNilImpl(tm, "match.PointerIsNil")
}
