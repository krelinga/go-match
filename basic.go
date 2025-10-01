package match

import (
	"cmp"

	"github.com/krelinga/go-typemap"
)

func Equal[T comparable](want T) Matcher[T] {
	tm := struct {
		typemap.StringFunc[T]
		typemap.DefaultCompare[T]
	}{
		StringFunc: DefaultString[T](),
	}
	return equalImpl(tm, "match.Equal", want)
}

func NotEqual[T comparable](other T) Matcher[T] {
	tm := struct {
		typemap.StringFunc[T]
		typemap.DefaultCompare[T]
	}{
		StringFunc: DefaultString[T](),
	}
	return notEqualImpl(tm, "match.NotEqual", other)
}

func LessThan[T cmp.Ordered](other T) Matcher[T] {
	tm := struct {
		typemap.StringFunc[T]
		typemap.DefaultOrder[T]
	}{
		StringFunc: DefaultString[T](),
	}
	return lessThanImpl(tm, "match.LessThan", other)
}

func LessThanOrEqual[T cmp.Ordered](other T) Matcher[T] {
	tm := struct {
		typemap.StringFunc[T]
		typemap.DefaultOrder[T]
	}{
		StringFunc: DefaultString[T](),
	}
	return lessThanOrEqualImpl(tm, "match.LessThanOrEqual", other)
}

func GreaterThan[T cmp.Ordered](other T) Matcher[T] {
	tm := struct {
		typemap.StringFunc[T]
		typemap.DefaultOrder[T]
	}{
		StringFunc: DefaultString[T](),
	}
	return greaterThanImpl(tm, "match.GreaterThan", other)
}

func GreaterThanOrEqual[T cmp.Ordered](other T) Matcher[T] {
	tm := struct {
		typemap.StringFunc[T]
		typemap.DefaultOrder[T]
	}{
		StringFunc: DefaultString[T](),
	}
	return greaterThanOrEqualImpl(tm, "match.GreaterThanOrEqual", other)
}
