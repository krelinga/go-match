package opts3

import (
	"errors"
	"reflect"
)

type Vals []reflect.Value

func NewVals(in ...any) Vals {
	vals := make(Vals, len(in))
	for i, v := range in {
		vals[i] = reflect.ValueOf(v)
	}
	return vals
}

func valueFor[T any](in T) reflect.Value {
	typ := reflect.TypeFor[T]()
	if typ.Kind() == reflect.Interface {
		v := reflect.New(typ).Elem()
		if reflect.ValueOf(in).IsValid() {
			v.Set(reflect.ValueOf(in))
		}
		return v
	}
	return reflect.ValueOf(in)
}

func NewVals1[T1 any](v T1) Vals {
	return Vals{valueFor(v)}
}

func NewVals2[T1, T2 any](v1 T1, v2 T2) Vals {
	return Vals{valueFor(v1), valueFor(v2)}
}

func NewVals3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Vals {
	return Vals{valueFor(v1), valueFor(v2), valueFor(v3)}
}

func zero[T any]() T {
	var v T
	return v
}

func wantImpl(vals Vals, types ...reflect.Type) ([]any, error) {
	if len(vals) != len(types) {
		return nil, errors.New("") // TODO
	}
	result := make([]any, len(types))
	for i, typ := range types {
		if !vals[i].IsValid() {
			return nil, errors.New("") // TODO
		}
		if vals[i].Type() != typ {
			return nil, errors.New("") // TODO
		}
		if !vals[i].CanInterface() {
			return nil, errors.New("") // TODO
		}
		result[i] = vals[i].Interface()
	}
	return result, nil
}

func Want1[T1 any](vals Vals) (T1, error) {
	asAny, err := wantImpl(vals, reflect.TypeFor[T1]())
	if err != nil {
		return zero[T1](), err
	}
	return asAny[0].(T1), nil
}

func Want2[T1, T2 any](vals Vals) (T1, T2, error) {
	asAny, err := wantImpl(vals, reflect.TypeFor[T1](), reflect.TypeFor[T2]())
	if err != nil {
		return zero[T1](), zero[T2](), err
	}
	return asAny[0].(T1), asAny[1].(T2), nil
}

func Want3[T1, T2, T3 any](vals Vals) (T1, T2, T3, error) {
	asAny, err := wantImpl(vals, reflect.TypeFor[T1](), reflect.TypeFor[T2](), reflect.TypeFor[T3]())
	if err != nil {
		return zero[T1](), zero[T2](), zero[T3](), err
	}
	return asAny[0].(T1), asAny[1].(T2), asAny[2].(T3), nil
}
