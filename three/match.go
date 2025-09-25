package match

import "fmt"

type Matcher[T any] interface {
	Match(got T) bool
	Condition(got T) string
}

type ParentMatcher[T any] interface {
	Matcher[T]
	Children(got T) Children
}

type Result struct {
	Condition string
	Children  Children
	Matched   bool
}

func NewResult[T any](got T, matcher Matcher[T]) Result {
	result := Result{
		Condition: matcher.Condition(got),
		Matched:   matcher.Match(got),
	}
	if pm, ok := matcher.(ParentMatcher[T]); ok {
		result.Children = pm.Children(got)
	}
	return result
}

type Child struct {
	Name   string
	Result Result
}

func NewChild[T any](name string, got T, matcher Matcher[T]) Child {
	return Child{
		Name:   name,
		Result: NewResult(got, matcher),
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

func Equal[T comparable](want T) Matcher[T] {
	return equalMatcher[T]{want: want}
}

type equalMatcher[T comparable] struct {
	want T
}

func (em equalMatcher[T]) Match(got T) bool {
	return got == em.want
}

func (em equalMatcher[T]) Condition(got T) string {
	return fmt.Sprintf("%v == %v", got, em.want)
}

func NotEqual[T comparable](want T) Matcher[T] {
	return notEqualMatcher[T]{want: want}
}

type notEqualMatcher[T comparable] struct {
	want T
}

func (nem notEqualMatcher[T]) Match(got T) bool {
	return got != nem.want
}

func (nem notEqualMatcher[T]) Condition(got T) string {
	return fmt.Sprintf("%v != %v", got, nem.want)
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
		children[i] = NewChild(fmt.Sprintf("child %d", i), got, matcher)
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
		children[i] = NewChild(fmt.Sprintf("child %d", i), got, matcher)
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
		Direct: []Child{NewChild("dereferenced", *got, dm.matcher)},
	}
}

func PointerEqual[T comparable](want *T) Matcher[*T] {
	return pointerEqualMatcher[T]{want: want}
}

type pointerEqualMatcher[T comparable] struct {
	want *T
}

func (pem pointerEqualMatcher[T]) Match(got *T) bool {
	if pem.want == nil || got == nil {
		return pem.want == got
	}
	return *pem.want == *got
}

func (pem pointerEqualMatcher[T]) Condition(_ *T) string {
	if pem.want == nil {
		return "== nil"
	}
	return fmt.Sprintf("== %v", *pem.want)
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
