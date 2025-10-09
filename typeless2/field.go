package typeless2

import (
	"fmt"
	"reflect"
)

func Field(name string, f Func) Func {
	return func(in Input) Output {
		if len(name) == 0 || (name[0] < 'A' || name[0] > 'Z') {
			return NewOutputErr(fmt.Errorf("field name %q must start with a capital letter", name))
		}
		got, err := Want1[any](in)
		if err != nil {
			return NewOutputErr(err)
		}
		gotVal := reflect.ValueOf(got)
		if !gotVal.IsValid() {
			return NewOutputErr(fmt.Errorf("%w: for field %s, wanted struct, got <nil>", ErrType, name))
		}
		if gotVal.Kind() != reflect.Struct {
			return NewOutputErr(fmt.Errorf("%w: for field %s, wanted struct, got %s", ErrType, name, gotVal.Type()))
		}
		gotField := gotVal.FieldByName(name)
		if !gotField.IsValid() {
			return NewOutputErr(fmt.Errorf("%w: struct %s has no field %s", ErrType, gotVal.Type(), name))
		}
		field := gotField.Interface()
		return f(NewInput(field))
	}
}