package opts

import (
	"errors"
	"fmt"
)

type Func func(opts Options, vals Values) (bool, error)

func Equal[T any](want T) Func {
	return func(opts Options, vals Values) (bool, error) {
		if len(vals) != 1 {
			return false, errors.New("Equal: expected exactly one argument")
		}
		val := vals[0]
		wantVal := ValueOf(want)
		if !val.IsValid() {
			return false, errors.New("Equal: invalid value")
		} else if val.Type() != wantVal.Type() {
			valStr, err := Fmt(opts, val)
			if err != nil {
				return false, err
			}
			return false, fmt.Errorf("Equal: type mismatch %s", valStr)
		} else if wantVal.Comparable() && val.Comparable() {
			if wantVal.Equal(val) {
				return true, nil
			} else {
				valStr, err := Fmt(opts, val)
				if err != nil {
					return false, err
				}
				return false, fmt.Errorf("Equal: value mismatch %v", valStr)
			}
		}
		return false, nil
	}
}

func WithOptions(opts ...Option) Bridge {
	return func(options Options, vals Values, m Func) (bool, error) {
		for _, o := range opts {
			options = o(options)
		}
		return m(options, vals)
	}
}
