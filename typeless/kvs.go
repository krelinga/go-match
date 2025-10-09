package typeless

import (
	"fmt"
	"reflect"

	"github.com/krelinga/go-match/matchfmt"
)

type KeyValue struct {
	Key   any
	Value any
}

type KeyValuesFunc func(got any) ([]KeyValue, error)

func DefaultKeyValues(got any) ([]KeyValue, error) {
	val := reflect.ValueOf(got)
	if !val.IsValid() {
		return nil, Error(ErrValue, "got value is invalid")
	}
	switch val.Kind() {
	case reflect.Map:
		result := make([]KeyValue, 0, val.Len())
		iter := val.MapRange()
		for iter.Next() {
			result = append(result, KeyValue{
				Key:   iter.Key().Interface(),
				Value: iter.Value().Interface(),
			})
		}
		return result, nil
	default:
		return nil, Error(ErrType, fmt.Sprintf("got value is not a map. type %q kind %s", val.Type(), val.Kind()))
	}
}

type KeyValuesMatcher struct {
	Inner Matcher
	KeyValues KeyValuesFunc
}

func KeyValues(matcher Matcher) KeyValuesMatcher {
	return KeyValuesMatcher{Inner: matcher, KeyValues: DefaultKeyValues}
}

func (m KeyValuesMatcher) Match(got any) (Matched, Explanation, error) {
	keyValues, err := m.KeyValues(got)
	if err != nil {
		return false, "", err
	}
	matched, explanation, err := m.Inner.Match(keyValues)
	if err != nil {
		return false, "", err
	}
	explanation = Explanation(matchfmt.Explain(bool(matched), "match.KeyValues", string(explanation)))
	return matched, explanation, nil
}