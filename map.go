package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
)

func NewMapHas[K comparable, V any](key K, val Matcher[V]) MapHas[K, V] {
	return MapHas[K, V]{Key: key, Val: val}
}

type MapHas[K comparable, V any] struct {
	Key K
	Val Matcher[V]
	KeyFormat func(t K) string
}

func (me MapHas[K, V]) Match(got map[K]V) bool {
	val, ok := got[me.Key]
	if !ok {
		return false
	}
	return me.Val.Match(val)
}

func (me MapHas[K, V]) Explain(got map[K]V) string {
	match := me.Match(got)
	var details []string
	keyStr := matchutil.FormatWith(me.Key, me.KeyFormat)
	expected := fmt.Sprintf("map has key %s with matching value", keyStr)
	if match {
		details = append(details, expected)
	} else {
		if val, ok := got[me.Key]; ok {
			actual := fmt.Sprintf("map has key %s with value that does not match", keyStr)
			details = append(details, matchutil.ActualVsExpected(actual, expected), Explain(val, me.Val))
		} else {
			actual := fmt.Sprintf("map does not have key %s", keyStr)
			details = append(details, matchutil.ActualVsExpected(actual, expected))
		}
	}
	return matchutil.Explain(match, matchutil.TypeName(me), details...)
}
