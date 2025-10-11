package opts2

import "reflect"

type cmpTagType struct{}

var cmpTag cmpTagType

type CmpFunc func(got, want reflect.Value) (bool, error)

func defaultCmpFunc(got, want reflect.Value) (bool, error) {
	if !got.IsValid() || !want.IsValid() {
		return false, ErrInvalid
	}
	if got.Type() != want.Type() {
		return false, ErrTypeMismatch
	}
	if !got.Comparable() || !want.Comparable() {
		return false, ErrInvalid
	}
	return got.Equal(want), nil
}

func Cmp(o Opts, got, want reflect.Value) (bool, error) {
	if !want.IsValid() {
		return false, ErrInvalid
	}
	var cmpFunc CmpFunc
	if a := o.Get(cmpTag, want.Type()); a != nil && a.(CmpFunc) != nil {
		cmpFunc = a.(CmpFunc)
	} else {
		cmpFunc = defaultCmpFunc
	}
	return cmpFunc(got, want)
}

func WithCmpFunc[T any](cmpFunc CmpFunc) Opt {
	return func(o Opts) Opts {
		return o.With(cmpTag, reflect.TypeFor[T](), cmpFunc)
	}
}

func WithCmpFuncAll(cmpFunc CmpFunc) Opt {
	return func(o Opts) Opts {
		return o.With(cmpTag, nil, cmpFunc)
	}
}

func EqualVals(want Vals) Func {
	return func(o Opts, got Vals) Out {
		if len(got) != len(want) {
			return Out{Err: ErrValCount}
		}
		matched := true
		for i := range got {
			if equal, err := Cmp(o, got[i], want[i]); err != nil {
				return Out{Err: err}
			} else if !equal {
				matched = false
			}
		}
		return Out{Matched: matched}
	}
}

func Equal(want ...any) Func {
	return EqualVals(NewVals(want...))
}
