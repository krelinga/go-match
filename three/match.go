package match

import "fmt"

type Matcher[T any] interface {
	Match(got T) bool
	Condition(got T) string
}

type Parent[T any] interface {
	Children(got T) Children
}

type Result struct {
	Condition string
	Children  Children
	Matched   bool
}

func Match[T any](got T, matcher Matcher[T]) bool {
	return matcher.Match(got)
}

func MatchResult[T any](got T, matcher Matcher[T]) Result {
	result := Result{
		Condition: matcher.Condition(got),
		Matched:   matcher.Match(got),
	}
	if pm, ok := matcher.(Parent[T]); ok {
		result.Children = pm.Children(got)
	}
	return result
}

type Child struct {
	Name   string
	Result Result
}

func MatchChild[T any](name string, got T, matcher Matcher[T]) Child {
	return Child{
		Name:   name,
		Result: MatchResult(got, matcher),
	}
}

type Children struct {
	Direct []Child
	Nested []NestedChildren
}

type NestedChildren struct {
	Category string
	Children
}

type Format[T any] func(got T) string

func (ff Format[T]) Format(in T) string {
	if ff != nil {
		return ff(in)
	}
	return fmt.Sprintf("%v", in)
}

func Equal[T comparable](want T) *EqualMatcher[T] {
	return &EqualMatcher[T]{want: want}
}

type EqualMatcher[T comparable] struct {
	want T
	f    Format[T]
}

func (em *EqualMatcher[T]) Match(got T) bool {
	return got == em.want
}

func (em *EqualMatcher[T]) Condition(got T) string {
	return fmt.Sprintf("%s == %s", em.f.Format(got), em.f.Format(em.want))
}

func (em *EqualMatcher[T]) WithFormat(f Format[T]) *EqualMatcher[T] {
	em.f = f
	return em
}

func NotEqual[T comparable](want T) *NotEqualMatcher[T] {
	return &NotEqualMatcher[T]{want: want}
}

type NotEqualMatcher[T comparable] struct {
	want T
	f    Format[T]
}

func (nem *NotEqualMatcher[T]) Match(got T) bool {
	return got != nem.want
}

func (nem *NotEqualMatcher[T]) Condition(got T) string {
	return fmt.Sprintf("%s != %s", nem.f.Format(got), nem.f.Format(nem.want))
}

func (nem *NotEqualMatcher[T]) WithFormat(f Format[T]) *NotEqualMatcher[T] {
	nem.f = f
	return nem
}

func AllOf[T any](matchers ...Matcher[T]) Matcher[T] {
	return allOfMatcher[T]{matchers: matchers}
}

type allOfMatcher[T any] struct {
	matchers []Matcher[T]
}

func (aom allOfMatcher[T]) Match(got T) bool {
	for _, matcher := range aom.matchers {
		if !matcher.Match(got) {
			return false
		}
	}
	return true
}

func (aom allOfMatcher[T]) Condition(_ T) string {
	return "all children match"
}

func (aom allOfMatcher[T]) Children(got T) Children {
	children := make([]Child, len(aom.matchers))
	for i, matcher := range aom.matchers {
		children[i] = MatchChild(fmt.Sprintf("child %d", i), got, matcher)
	}
	return Children{Direct: children}
}

func AnyOf[T any](matchers ...Matcher[T]) Matcher[T] {
	return anyOfMatcher[T]{matchers: matchers}
}

type anyOfMatcher[T any] struct {
	matchers []Matcher[T]
}

func (aom anyOfMatcher[T]) Match(got T) bool {
	for _, matcher := range aom.matchers {
		if matcher.Match(got) {
			return true
		}
	}
	return false
}

func (aom anyOfMatcher[T]) Condition(_ T) string {
	return "any child matches"
}

func (aom anyOfMatcher[T]) Children(got T) Children {
	children := make([]Child, len(aom.matchers))
	for i, matcher := range aom.matchers {
		children[i] = MatchChild(fmt.Sprintf("child %d", i), got, matcher)
	}
	return Children{Direct: children}
}

func Deref[T any](matcher Matcher[T]) Matcher[*T] {
	return derefMatcher[T]{matcher: matcher}
}

type derefMatcher[T any] struct {
	matcher Matcher[T]
}

func (dm derefMatcher[T]) Match(got *T) bool {
	if got == nil {
		return false
	}
	return dm.matcher.Match(*got)
}

func (dm derefMatcher[T]) Condition(_ *T) string {
	return "dereferenced pointer matches"
}

func (dm derefMatcher[T]) Children(got *T) Children {
	if got == nil {
		return Children{}
	}
	return Children{
		Direct: []Child{MatchChild("dereferenced", *got, dm.matcher)},
	}
}

func PointerEqual[T comparable](want *T) *PointerEqualMatcher[T] {
	return &PointerEqualMatcher[T]{want: want}
}

type PointerEqualMatcher[T comparable] struct {
	want *T
	f    Format[T]
}

func (pem *PointerEqualMatcher[T]) Match(got *T) bool {
	if pem.want == nil || got == nil {
		return pem.want == got
	}
	return *pem.want == *got
}

func (pem *PointerEqualMatcher[T]) Condition(_ *T) string {
	if pem.want == nil {
		return "== nil"
	}
	return fmt.Sprintf("== %v", pem.f.Format(*pem.want))
}

func (pem *PointerEqualMatcher[T]) WithFormat(f Format[T]) *PointerEqualMatcher[T] {
	pem.f = f
	return pem
}

func Elements[T any](matchers ...Matcher[T]) *ElementsMatcher[T] {
	return &ElementsMatcher[T]{matchers: matchers}
}

type ElementsMatcher[T any] struct {
	matchers  []Matcher[T]
	unordered bool
}

func (em *ElementsMatcher[T]) Unordered(u bool) *ElementsMatcher[T] {
	em.unordered = u
	return em
}

func (em *ElementsMatcher[T]) Match(got []T) bool {
	if em.matchers == nil {
		panic("ElementsMatcher not initialized.  Use Elements(...) to create one.")
	}
	if len(got) != len(em.matchers) {
		return false
	}
	if em.unordered {
		return em.matchUnordered(got)
	}
	return em.matchOrdered(got)
}

// TODO: support Children() on ElementsMatcher.

func (em *ElementsMatcher[T]) matchOrdered(got []T) bool {
	for i, matcher := range em.matchers {
		if !matcher.Match(got[i]) {
			return false
		}
	}
	return true
}

func (em *ElementsMatcher[T]) matchUnordered(got []T) bool {
	used := make([]bool, len(got))
	for _, matcher := range em.matchers {
		matched := false
		for i, g := range got {
			if !used[i] && matcher.Match(g) {
				used[i] = true
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

func (em *ElementsMatcher[T]) Condition(_ []T) string {
	if em.unordered {
		return "elements match (unordered)"
	}
	return "elements match (ordered)"
}

func Slice[T any](want []T, factory func(t T) Matcher[T]) Matcher[[]T] {
	matchers := make([]Matcher[T], len(want))
	for i, w := range want {
		matchers[i] = factory(w)
	}
	return Elements(matchers...)
}
