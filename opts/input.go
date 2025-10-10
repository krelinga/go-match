package opts

import "reflect"

type Values []reflect.Value

func NewValues1[T1 any](v1 T1) Values {
	return Values{ValueOf(v1)}
}

func NewValues2[T1, T2 any](v1 T1, v2 T2) Values {
	return Values{ValueOf(v1), ValueOf(v2)}
}

func NewValues3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Values {
	return Values{ValueOf(v1), ValueOf(v2), ValueOf(v3)}
}