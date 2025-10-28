package opts3

import (
	"errors"
	"reflect"
)

type cmpTag struct{}

func Cmp(env Env, a, b reflect.Value) (bool, error) {
	if !a.IsValid() || !b.IsValid() {
		return false, errors.New("") // TODO
	}
	if a.Type() != b.Type() {
		return false, errors.New("") // TODO
	}
	typ := a.Type()
	if !a.CanInterface() || !b.CanInterface() {
		return false, errors.New("") // TODO
	}
	aAny, bAny := a.Interface(), b.Interface()
	if f, ok := env.Get(cmpTag{}, typ); ok && f != nil {
		return f.(func(any, any) bool)(aAny, bAny), nil
	} else {
		return DefaultCmp(aAny, bAny), nil
	}
}

func DefaultCmp(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

func CmpOpt(typ reflect.Type, f func(any, any) bool) Opt {
	return OptFunc(func(env Env) Env {
		return WrapEnv(env, cmpTag{}, typ, f)
	})
}

func CmpOptFor[T any](f func(T, T) bool) Opt {
	typ := reflect.TypeFor[T]()
	return CmpOpt(typ, func(a, b any) bool {
		return f(a.(T), b.(T))
	})
}

func CmpOptAll(f func(any, any) bool) Opt {
	return CmpOpt(nil, f)
}