package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
	"github.com/krelinga/go-typemap"
)

func hasKeyImpl[T, K any](containerTm typemap.HasKey[T, K], keyTm typemap.String[K], name string, key K) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = containerTm.HasKey(got, key)
		expected := fmt.Sprintf("has key %s", keyTm.String(key))
		if matched {
			explanation = expected
		} else {
			actual := fmt.Sprintf("key %s not found", keyTm.String(key))
			explanation = matchutil.Explain(matched, name, matchutil.ActualVsExpected(actual, expected))
		}
		return
	})
}

func HasKeyTm[T, K any](containerTm typemap.HasKey[T, K], keyTm typemap.String[K], key K) Matcher[T] {
	return hasKeyImpl(containerTm, keyTm, "match.HasKeyTm", key)
}

func StringHasIndex[T ~string](index int) Matcher[T] {
	contTm := typemap.ForStringLike[T]{}
	keyTm := typemap.ForInt{
		StringFunc: DefaultString[int](),
	}
	return hasKeyImpl(contTm, keyTm, "match.StringHasIndex", index)
}

func SliceHasIndex[T ~[]E, E any](index int) Matcher[T] {
	contTm := typemap.ForSliceLike[T, E]{}
	keyTm := typemap.ForInt{
		StringFunc: DefaultString[int](),
	}
	return hasKeyImpl(contTm, keyTm, "match.SliceHasIndex", index)
}

func MapHasKey[T ~map[K]V, K comparable, V any](key K) Matcher[T] {
	contTm := typemap.ForMapLike[T, K, V]{}
	keyTm := struct {
		typemap.StringFunc[K]
	}{
		StringFunc: DefaultString[K](),
	}
	return hasKeyImpl(contTm, keyTm, "match.MapHasKey", key)
}
