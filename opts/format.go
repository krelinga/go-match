package opts

import (
	"errors"
	"fmt"
	"reflect"
)

type FmtFunc func(reflect.Value) (string, error)

type fmtOpKey struct{}

var fmtKey = fmtOpKey{}

func SetFmtFunc(typeKey reflect.Type, fmt FmtFunc) Option {
	return func(options Options) Options {
		return options.With(typeKey, fmtKey, fmt)
	}
}

func Fmt(options Options, val reflect.Value) (string, error) {
	fmtFunc, ok := options.Get(val.Type(), fmtKey).(FmtFunc)
	if ok && fmtFunc != nil {
		return fmtFunc(val)
	}
	if !val.IsValid() {
		return "", errors.New("invalid value")
	} else if !val.CanInterface() {
		return "", errors.New("cannot interface value")
	} else if val.Kind() == reflect.Ptr && !val.IsNil() {
		str, err := Fmt(options, val.Elem())
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("PointerTo(%s)", str), nil
	}
	return fmt.Sprintf("%#v", val.Interface()), nil
}
