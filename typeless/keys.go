package typeless

import (
	"fmt"
	"reflect"

	"github.com/krelinga/go-match/matchfmt"
)

type KeysFunc func(got any) ([]any, error)

func DefaultKeys(got any) ([]any, error) {
	val := reflect.ValueOf(got)
	if !val.IsValid() {
		return nil, Error(ErrValue, "got value is invalid")
	}
	switch val.Kind() {
	case reflect.Map:
		keys := val.MapKeys()
		result := make([]any, len(keys))
		for i, key := range keys {
			result[i] = key.Interface()
		}
		return result, nil
	default:
		return nil, Error(ErrType, fmt.Sprintf("got value is not a map. type %q kind %s", val.Type(), val.Kind()))
	}
}

type KeysMatcher struct {
	Inner Matcher
	Keys  KeysFunc
}

func Keys(matcher Matcher) KeysMatcher {
	return KeysMatcher{Inner: matcher, Keys: DefaultKeys}
}

func (m KeysMatcher) Match(got any) (Matched, Explanation, error) {
	keys, err := m.Keys(got)
	if err != nil {
		return false, "", err
	}
	matched, explanation, err := m.Inner.Match(keys)
	if err != nil {
		return false, "", err
	}
	explanation = Explanation(matchfmt.Explain(bool(matched), "match.Keys", string(explanation)))
	return matched, explanation, nil
}