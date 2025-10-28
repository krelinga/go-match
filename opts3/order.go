package opts3

import "reflect"

type orderTag struct{}

func Order(env Env, a, b reflect.Value) (int, error) {
	if !a.IsValid() || !b.IsValid() {
		return 0, nil // TODO
	}
	if a.Type() != b.Type() {
		return 0, nil // TODO
	}
	typ := a.Type()
	if !a.CanInterface() || !b.CanInterface() {
		return 0, nil // TODO
	}
	aAny, bAny := a.Interface(), b.Interface()
	if f, ok := env.Get(orderTag{}, typ); ok && f != nil {
		return f.(func(any, any) int)(aAny, bAny), nil
	} else {
		return DefaultOrder(aAny, bAny), nil
	}
}

// TODO: this will require reflection primitives ... rethink the api.
func DefaultOrder(a, b any) int {
	return 0 // TODO
}
