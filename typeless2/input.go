package typeless2

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrInputLen = errors.New("input length mismatch")
var ErrType = errors.New("type mismatch")

type Input []any

func NewInput(vals ...any) Input {
	return Input(vals)
}

func zero[T any]() T {
	var t T
	return t
}

func wantTypeError[T any](pos int, got any) error {
	var gotStr string
	gotVal := reflect.ValueOf(got)
	if !gotVal.IsValid() {
		gotStr = "<nil>"
	} else {
		gotStr = gotVal.Type().String()
	}

	wantStr := reflect.TypeFor[T]().String()
	return fmt.Errorf("%w: for arg %d, wanted %s, got %s", ErrType, pos, wantStr, gotStr)
}

func wantLenError(want, got int) error {
	return fmt.Errorf("%w: wanted %d, got %d", ErrInputLen, want, got)
}

func Want1[T1 any](i Input) (T1, error) {
	if len(i) != 1 {
		return zero[T1](), wantLenError(1, len(i))
	}
	v1, ok := i[0].(T1)
	if !ok {
		return zero[T1](), wantTypeError[T1](1, i[0])
	}
	return v1, nil
}

func Want2[T1, T2 any](i Input) (T1, T2, error) {
	if len(i) != 2 {
		return zero[T1](), zero[T2](), wantLenError(2, len(i))
	}
	v1, ok := i[0].(T1)
	if !ok {
		return zero[T1](), zero[T2](), wantTypeError[T1](1, i[0])
	}
	v2, ok := i[1].(T2)
	if !ok {
		return zero[T1](), zero[T2](), wantTypeError[T2](2, i[1])
	}
	return v1, v2, nil
}

// TODO: more WantN functions?