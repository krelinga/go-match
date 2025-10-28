package opts3

import (
	"fmt"
	"reflect"
)

func Fmt(env Env, val reflect.Value) string {
	if !val.IsValid() {
		return "<invalid>"
	}
	if !val.CanInterface() {
		return "<uninterfaceable>"
	}
	typ := val.Type()
	if f, ok := env.Get(fmtTag{}, typ); ok && f != nil {
		return f.(func(any) string)(val.Interface())
	} else {
		return DefaultFmt(val.Interface())
	}
}

func DefaultFmt(v any) string {
	return fmt.Sprintf("%#v", v)
}

type fmtTag struct{}

func FmtOpt(typ reflect.Type, f func(any) string) Opt {
	return OptFunc(func(env Env) Env {
		return WrapEnv(env, fmtTag{}, typ, f)
	})
}

func FmtOptFor[T any](f func(T) string) Opt {
	typ := reflect.TypeFor[T]()
	return FmtOpt(typ, func(v any) string {
		return f(v.(T))
	})
}

func FmtOptAll(f func(any) string) Opt {
	return FmtOpt(nil, f)
}
