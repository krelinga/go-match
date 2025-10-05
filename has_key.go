package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
	"github.com/krelinga/go-typemap"
)

func hasKeyImpl[T, K any](containerTm typemap.HasKey[T, K], keyTm typemap.String[K], matcherName, keyName string, key K) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = containerTm.HasKey(got, key)
		expected := fmt.Sprintf("has %s %s", keyName, keyTm.String(key))
		var detail string
		if matched {
			detail = expected
		} else {
			actual := fmt.Sprintf("%s %s not found", keyName, keyTm.String(key))
			detail = matchutil.ActualVsExpected(actual, expected)
		}
		explanation = matchutil.Explain(matched, matcherName, detail)
		return
	})
}

func HasKeyTm[T, K any](containerTm typemap.HasKey[T, K], keyTm typemap.String[K], key K) Matcher[T] {
	return hasKeyImpl(containerTm, keyTm, "match.HasKeyTm", "key", key)
}

func StringLikeHasIndex[T ~string](index int) Matcher[T] {
	contTm := typemap.ForStringLike[T]{}
	keyTm := typemap.ForInt{
		StringFunc: DefaultString[int],
	}
	return hasKeyImpl(contTm, keyTm, "match.StringLikeHasIndex", "index", index)
}

func StringHasIndex(index int) Matcher[string] {
	contTm := typemap.ForString{}
	keyTm := typemap.ForInt{
		StringFunc: DefaultString[int],
	}
	return hasKeyImpl(contTm, keyTm, "match.StringHasIndex", "index", index)
}

func SliceLikeHasIndex[T ~[]E, E any](index int) Matcher[T] {
	contTm := typemap.ForSliceLike[T, E]{}
	keyTm := typemap.ForInt{
		StringFunc: DefaultString[int],
	}
	return hasKeyImpl(contTm, keyTm, "match.SliceLikeHasIndex", "index", index)
}

func SliceHasIndex[E any](index int) Matcher[[]E] {
	contTm := typemap.ForSlice[E]{}
	keyTm := typemap.ForInt{
		StringFunc: DefaultString[int],
	}
	return hasKeyImpl(contTm, keyTm, "match.SliceHasIndex", "index", index)
}

func MapLikeHasKey[T ~map[K]V, K comparable, V any](key K) Matcher[T] {
	contTm := typemap.ForMapLike[T, K, V]{}
	keyTm := struct {
		typemap.StringFunc[K]
	}{
		StringFunc: DefaultString[K],
	}
	return hasKeyImpl(contTm, keyTm, "match.MapLikeHasKey", "key", key)
}

func MapHasKey[K comparable, V any](key K) Matcher[map[K]V] {
	contTm := typemap.ForMap[K, V]{}
	keyTm := struct {
		typemap.StringFunc[K]
	}{
		StringFunc: DefaultString[K],
	}
	return hasKeyImpl(contTm, keyTm, "match.MapHasKey", "key", key)
}
