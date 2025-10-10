package opts

import "reflect"

func ValueOf[T any](v T) reflect.Value {
	val := reflect.New(reflect.TypeFor[T]()).Elem()
	val.Set(reflect.ValueOf(v))
	return val
}