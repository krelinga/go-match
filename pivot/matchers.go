package pivot

import (
	"errors"
	"reflect"
)

var ErrTypeMismatch = errors.New("type mismatch")

type Equal struct {
	X ValueEqualer
}

func (e Equal) Match(got any) (bool, error) {
	if !e.X.SameTypeAs(got) {
		return false, ErrTypeMismatch
	}
	return e.X.Equal(got), nil
}

type NotEqual struct {
	X ValueEqualer
}

func (n NotEqual) Match(got any) (bool, error) {
	if !n.X.SameTypeAs(got) {
		return false, ErrTypeMismatch
	}
	return !n.X.Equal(got), nil
}

type LessThan struct {
	X ValueComparer
}

func (lt LessThan) Match(got any) (bool, error) {
	if !lt.X.SameTypeAs(got) {
		return false, ErrTypeMismatch
	}
	return lt.X.Compare(got) < 0, nil
}

type LessThanOrEqual struct {
	X ValueComparer
}

func (lte LessThanOrEqual) Match(got any) (bool, error) {
	if !lte.X.SameTypeAs(got) {
		return false, ErrTypeMismatch
	}
	return lte.X.Compare(got) <= 0, nil
}

type GreaterThan struct {
	X ValueComparer
}

func (gt GreaterThan) Match(got any) (bool, error) {
	if !gt.X.SameTypeAs(got) {
		return false, ErrTypeMismatch
	}
	return gt.X.Compare(got) > 0, nil
}

type GreaterThanOrEqual struct {
	X ValueComparer
}

func (gte GreaterThanOrEqual) Match(got any) (bool, error) {
	if !gte.X.SameTypeAs(got) {
		return false, ErrTypeMismatch
	}
	return gte.X.Compare(got) >= 0, nil
}

type AllOf struct {
	M []Matcher
}

func (a AllOf) Match(got any) (bool, error) {
	for _, m := range a.M {
		if ok, err := m.Match(got); !ok {
			return false, err
		}
	}
	return true, nil
}

type Lengther interface {
	Length() int
}

func newBuiltinLengther(v any) (Lengther, bool) {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return builtinLengther(val), true
	default:
		return nil, false
	}
}

type builtinLengther reflect.Value

func (bl builtinLengther) Length() int {
	return reflect.Value(bl).Len()
}

type Length struct {
	M Matcher
}

func (l Length) Match(got any) (bool, error) {
	if lt, ok := got.(Lengther); ok {
		return l.M.Match(lt.Length())
	} else if bl, ok := newBuiltinLengther(got); ok {
		return l.M.Match(bl.Length())
	}
	return false, ErrTypeMismatch
}

type HasKeyer interface {
	HasKey(key any) (bool, error)
}

func newBuiltinHasKeyer(v any) (HasKeyer, bool) {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array:
		return builtinHasKeyer(val), true
	default:
		return nil, false
	}
}

type builtinHasKeyer reflect.Value

func (bhk builtinHasKeyer) HasKey(key any) (bool, error) {
	switch reflect.Value(bhk).Kind() {
	case reflect.Map:
		return bhk.mapHasKey(key)
	case reflect.Slice, reflect.Array:
		return bhk.sliceHasKey(key)
	default:
		panic("unreachable")
	}
}

func (bhk builtinHasKeyer) mapHasKey(key any) (bool, error) {
	keyVal := reflect.ValueOf(key)
	mapVal := reflect.Value(bhk)
	if !keyVal.Type().AssignableTo(mapVal.Type().Key()) {
		return false, ErrTypeMismatch
	}
	return mapVal.MapIndex(keyVal).IsValid(), nil
}

func (bhk builtinHasKeyer) sliceHasKey(key any) (bool, error) {
	index, ok := key.(int)
	if !ok {
		return false, ErrTypeMismatch
	}
	sliceVal := reflect.Value(bhk)
	return index >= 0 && index < sliceVal.Len(), nil
}

type HasKey[K comparable] struct {
	Key K
}

func (hk HasKey[K]) Match(got any) (bool, error) {
	if hki, ok := got.(HasKeyer); ok {
		return hki.HasKey(hk.Key)
	} else if bhk, ok := newBuiltinHasKeyer(got); ok {
		return bhk.HasKey(hk.Key)
	}
	return false, ErrTypeMismatch
}

type HasValer interface {
	HasVal(val any) (bool, error)
}

type HasVal struct {
	M Matcher
}

func (hv HasVal) Match(got any) (bool, error) {
	return false, nil // TODO
}
