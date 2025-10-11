package opts2

import (
	"errors"
	"reflect"
)

type Vals []reflect.Value

func NewVals(vs ...any) Vals {
	vals := make(Vals, len(vs))
	for i, v := range vs {
		vals[i] = reflect.ValueOf(v)
	}
	return vals
}

func typedVal[T any](v T) reflect.Value {
	x := reflect.New(reflect.TypeFor[T]()).Elem()
	x.Set(reflect.ValueOf(v))
	return x
}

func NewVals1[T1 any](v1 T1) Vals {
	return Vals{typedVal(v1)}
}

func NewVals2[T1, T2 any](v1 T1, v2 T2) Vals {
	return Vals{typedVal(v1), typedVal(v2)}
}

func NewVals3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Vals {
	return Vals{typedVal(v1), typedVal(v2), typedVal(v3)}
}

func NewVals4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Vals {
	return Vals{typedVal(v1), typedVal(v2), typedVal(v3), typedVal(v4)}
}

func zero[T any]() T {
	var t T
	return t
}

var (
	ErrValCount     = errors.New("wrong number of vals")
	ErrTypeMismatch = errors.New("type mismatch")
	ErrInvalid      = errors.New("value is invalid")
	ErrInterface    = errors.New("value is not interfaceable")
)

func want[T any](v reflect.Value) (T, error) {
	if v.Type() != reflect.TypeFor[T]() {
		return zero[T](), ErrTypeMismatch
	}
	if !v.IsValid() {
		return zero[T](), ErrInvalid
	}
	if !v.CanInterface() {
		return zero[T](), ErrInterface
	}
	return v.Interface().(T), nil
}

func Want1[T0 any](v Vals) (T0, error) {
	if len(v) != 1 {
		return zero[T0](), ErrValCount
	}
	return want[T0](v[0])
}

func Want1Val(v Vals) (reflect.Value, error) {
	if len(v) != 1 {
		return reflect.Value{}, ErrValCount
	}
	return v[0], nil
}

func Want2[T0, T1 any](v Vals) (T0, T1, error) {
	if len(v) != 2 {
		return zero[T0](), zero[T1](), ErrValCount
	}
	v0, err := want[T0](v[0])
	if err != nil {
		return zero[T0](), zero[T1](), err
	}
	v1, err := want[T1](v[1])
	if err != nil {
		return zero[T0](), zero[T1](), err
	}
	return v0, v1, nil
}

func Want2Val(v Vals) (reflect.Value, reflect.Value, error) {
	if len(v) != 2 {
		return reflect.Value{}, reflect.Value{}, ErrValCount
	}
	return v[0], v[1], nil
}

func Want3[T0, T1, T2 any](v Vals) (T0, T1, T2, error) {
	if len(v) != 3 {
		return zero[T0](), zero[T1](), zero[T2](), ErrValCount
	}
	v0, err := want[T0](v[0])
	if err != nil {
		return zero[T0](), zero[T1](), zero[T2](), err
	}
	v1, err := want[T1](v[1])
	if err != nil {
		return zero[T0](), zero[T1](), zero[T2](), err
	}
	v2, err := want[T2](v[2])
	if err != nil {
		return zero[T0](), zero[T1](), zero[T2](), err
	}
	return v0, v1, v2, nil
}

func Want3Val(v Vals) (reflect.Value, reflect.Value, reflect.Value, error) {
	if len(v) != 3 {
		return reflect.Value{}, reflect.Value{}, reflect.Value{}, ErrValCount
	}
	return v[0], v[1], v[2], nil
}

func Want4[T0, T1, T2, T3 any](v Vals) (T0, T1, T2, T3, error) {
	if len(v) != 4 {
		return zero[T0](), zero[T1](), zero[T2](), zero[T3](), ErrValCount
	}
	v0, err := want[T0](v[0])
	if err != nil {
		return zero[T0](), zero[T1](), zero[T2](), zero[T3](), err
	}
	v1, err := want[T1](v[1])
	if err != nil {
		return zero[T0](), zero[T1](), zero[T2](), zero[T3](), err
	}
	v2, err := want[T2](v[2])
	if err != nil {
		return zero[T0](), zero[T1](), zero[T2](), zero[T3](), err
	}
	v3, err := want[T3](v[3])
	if err != nil {
		return zero[T0](), zero[T1](), zero[T2](), zero[T3](), err
	}
	return v0, v1, v2, v3, nil
}

func Want4Val(v Vals) (reflect.Value, reflect.Value, reflect.Value, reflect.Value, error) {
	if len(v) != 4 {
		return reflect.Value{}, reflect.Value{}, reflect.Value{}, reflect.Value{}, ErrValCount
	}
	return v[0], v[1], v[2], v[3], nil
}
