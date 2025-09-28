package match

type Matcher[T any] interface {
	Match(got T) bool
}

type Explainer[T any] interface {
	Explain(got T) string
}

func matchEmoji(matched bool) string {
	if matched {
		return "✅"
	}
	return "❌"
}

func Match[T any](got T, matcher Matcher[T]) bool {
	return matcher.Match(got)
}

func Explain[T any](got T, matcher Matcher[T]) string {
	return "" // TODO
}

func NewEqual[T comparable](x T) Equal[T] {
	return Equal[T]{X: x}
}

type Equal[T comparable] struct {
	X T
	Format func(t T) string
}

func (e Equal[T]) Match(got T) bool {
	return got == e.X
}

func (e Equal[T]) Explain(got T) string {
	return "" // TODO
}

func NewNotEqual[T comparable](x T) NotEqual[T] {
	return NotEqual[T]{X: x}
}

type NotEqual[T comparable] struct {
	X T
	Format func(t T) string
}

func (ne NotEqual[T]) Match(got T) bool {
	return got != ne.X
}

func (ne NotEqual[T]) Explain(got T) string {
	return "" // TODO
}

func NewAllOf[T any](m ...Matcher[T]) AllOf[T] {
	return AllOf[T]{M: m}
}

type AllOf[T any] struct {
	M []Matcher[T]
}

func (a AllOf[T]) Match(got T) bool {
	for _, m := range a.M {
		if !m.Match(got) {
			return false
		}
	}
	return true
}

func (a AllOf[T]) Explain(_ T) string {
	return "" // TODO
}

type AnyOf[T any] struct {
	M []Matcher[T]
}

func NewAnyOf[T any](m ...Matcher[T]) AnyOf[T] {
	return AnyOf[T]{M: m}
}

func (a AnyOf[T]) Match(got T) bool {
	for _, m := range a.M {
		if m.Match(got) {
			return true
		}
	}
	return false
}

func (a AnyOf[T]) Explain(_ T) string {
	return "" // TODO
}

func NewWhenDeref[T any](m Matcher[T]) WhenDeref[T] {
	return WhenDeref[T]{M: m}
}

type WhenDeref[T any] struct {
	M Matcher[T]
}

func (p WhenDeref[T]) Match(got *T) bool {
	if got == nil {
		return false
	}
	return p.M.Match(*got)
}

func (p WhenDeref[T]) Explain(got *T) string {
	return "" // TODO
}

func NewSliceElems[T any](m ...Matcher[T]) SliceElems[T] {
	return SliceElems[T]{M: m}
}

type SliceElems[T any] struct {
	M []Matcher[T]
}

func (s SliceElems[T]) Match(got []T) bool {
	if len(got) != len(s.M) {
		return false
	}
	for i, elem := range got {
		if !s.M[i].Match(elem) {
			return false
		}
	}
	return true
}

func (s SliceElems[T]) Explain(_ []T) string {
	return "" // TODO
}
