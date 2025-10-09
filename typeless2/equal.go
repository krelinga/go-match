package typeless2

import (
	"fmt"
	"reflect"
)

type Cmp func(a, b any) (bool, error)

func cmpDefTypeError[T any](got any, label string) error {
	wantStr := reflect.TypeFor[T]().String()
	gotStr := reflect.TypeOf(got).String()
	return fmt.Errorf("%w: for %s: wanted %s, got %s", ErrType, label, wantStr, gotStr)
}

func CmpDefault(got, want any) (bool, error) {
	wantType := reflect.TypeOf(want)
	gotType := reflect.TypeOf(got)
	if wantType != gotType {
		return false, fmt.Errorf("%w: wanted %s, got %s", ErrType, wantType, gotType)
	}
	if !wantType.Comparable() {
		return false, fmt.Errorf("%w: type %s is not comparable", ErrType, wantType)
	}
	return reflect.ValueOf(got).Equal(reflect.ValueOf(want)), nil
}

type EqualOpt struct {
	Fmt Fmt
	Cmp Cmp
}

func Equal(want any) Func {
	return EqualO(want, EqualOpt{})
}

func EqualO(want any, opt EqualOpt) Func {
	if opt.Cmp == nil {
		opt.Cmp = CmpDefault
	}
	if opt.Fmt == nil {
		opt.Fmt = FmtDef
	}
	return func(in Input) Output {
		got, err := Want1[any](in)
		if err != nil {
			return NewOutputErr(err)
		}
		eq, err := opt.Cmp(got, want)
		if err != nil {
			return NewOutputErr(err)
		}
		if eq {
			return Output{Matched: true}
		} else {
			wantStr, err := opt.Fmt(want)
			if err != nil {
				return NewOutputErr(err)
			}
			gotStr, err := opt.Fmt(got)
			if err != nil {
				return NewOutputErr(err)
			}
			return Output{
				Matched:     false,
				Explanation: fmt.Sprintf("wanted %s, got %s", wantStr, gotStr),
			}
		}
	}
}
