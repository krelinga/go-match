package opts2

import "reflect"

func Field(name string) Bridge {
	return func(o Opts, vals Vals, f Func) Out {
		val, err := Want1Val(vals)
		if err != nil {
			return Out{Err: err}
		}
		if !val.IsValid() {
			return Out{Err: ErrInvalid}
		}
		if val.Kind() == reflect.Ptr && val.Type().Elem().Kind() == reflect.Struct {
			if val.IsNil() {
				return Out{Matched: false, Note: "nil value"}
			} else {
				val = val.Elem()
			}
		}
		if val.Kind() != reflect.Struct {
			return Out{Err: ErrTypeMismatch}
		}
		field := val.FieldByName(name)
		if !field.IsValid() {
			return Out{Err: ErrInvalid, Note: "no such field"}
		}
		return f(o, Vals{field})
	}
}