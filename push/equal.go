package push

type CompareFunc[T any] func(a, b T) bool

func DefaultCompare[T comparable]() CompareFunc[T] {
	return func(a, b T) bool {
		return a == b
	}
}

type EqualOpts[T any] struct {
	Want T
	Compare CompareFunc[T]
	Formatter Formatter[T]
}

func (o EqualOpts[T]) M(got T) Result {
	equal := o.Compare(got, o.Want)
	var reason string
	if !equal {
		reason = "not equal"
		if o.Formatter != nil {
			reason = "got " + o.Formatter.Format(got) + ", want " + o.Formatter.Format(o.Want)
		}
	}
	return Result{Matched: equal, Reason: reason}
}

func Equal[T comparable](want T) M[T] {
	return EqualOpts[T]{Want: want, Compare: DefaultCompare[T](), Formatter: DefaultFormatter[T]()}.M
}
