package match

import (
	"fmt"
	"reflect"
	"strings"
)

type Matcher[T any] interface {
	Match(got T) bool
}

type Explainer[T any] interface {
	Explain(matched bool, got T) string
}

type Unwrapper[T any] interface {
	Unwrap(got T) *ResultTree
}

type Result struct {
	MatcherType string
	Explanation string
	Unwrapped   *ResultTree
	Matched     bool
}

func (r Result) String() string {
	parts := []string{}
	parts = append(parts, fmt.Sprintf("%s %s", matchEmoji(r.Matched), r.MatcherType))
	if r.Explanation != "" {
		parts = append(parts, fmt.Sprintf("   %s", indent(r.Explanation, 1)))
	}
	if r.Unwrapped != nil {
		parts = append(parts, fmt.Sprintf("   %s", indent(r.Unwrapped.String(), 1)))
	}
	return strings.Join(parts, "\n")
}

func matchEmoji(matched bool) string {
	if matched {
		return "✅"
	}
	return "❌"
}

func indent(s string, depth int) string {
	padding := strings.Repeat("   ", depth)
	return strings.ReplaceAll(s, "\n", "\n"+padding)
}

func Match[T any](got T, matcher Matcher[T]) bool {
	return matcher.Match(got)
}

func MatchResult[T any](got T, matcher Matcher[T]) Result {
	result := Result{
		MatcherType: reflect.TypeOf(matcher).String(),
		Matched:     matcher.Match(got),
	}
	if cm, ok := matcher.(Explainer[T]); ok {
		result.Explanation = cm.Explain(result.Matched, got)
	}
	if pm, ok := matcher.(Unwrapper[T]); ok {
		result.Unwrapped = pm.Unwrap(got)
	}
	return result
}

type ResultTree struct {
	Root     []Result
	Branches []ResultBranch
}

func (rt *ResultTree) String() string {
	parts := []string{}
	for _, r := range rt.Root {
		parts = append(parts, r.String())
	}
	for _, b := range rt.Branches {
		parts = append(parts, fmt.Sprintf("%s:", b.Name), b.ResultTree.String())
	}
	return strings.Join(parts, "\n")
}

func NewResultTreeRoot(results ...Result) *ResultTree {
	return &ResultTree{
		Root: results,
	}
}

type ResultBranch struct {
	Name string
	ResultTree
}

type FormatHelper[T any] struct {
	ff func(t T) string
}

func (fh *FormatHelper[T]) Format(t T) string {
	if fh.ff != nil {
		return fh.ff(t)
	}
	return fmt.Sprintf("%v", t)
}

func (fh *FormatHelper[T]) Set(ff func(t T) string) {
	fh.ff = ff
}

func ActualVsExpected(matched bool, actual, expected string) string {
	if matched {
		return expected
	}
	sb := &strings.Builder{}
	sb.WriteString("ACTUAL  : ")
	sb.WriteString(actual)
	sb.WriteString("\n")
	sb.WriteString("EXPECTED: ")
	sb.WriteString(expected)
	return sb.String()
}

func Equal[T comparable](want T) *EqualMatcher[T] {
	return &EqualMatcher[T]{want: want}
}

type EqualMatcher[T comparable] struct {
	want T
	fh   FormatHelper[T]
}

func (em *EqualMatcher[T]) Match(got T) bool {
	return got == em.want
}

func (em *EqualMatcher[T]) Explain(matched bool, got T) string {
	actual := fmt.Sprintf("got = %s", em.fh.Format(got))
	expected := fmt.Sprintf("got = %s", em.fh.Format(em.want))
	return ActualVsExpected(matched, actual, expected)
}

func (em *EqualMatcher[T]) WithFormat(f func(t T) string) *EqualMatcher[T] {
	em.fh.Set(f)
	return em
}

func NotEqual[T comparable](want T) *NotEqualMatcher[T] {
	return &NotEqualMatcher[T]{want: want}
}

type NotEqualMatcher[T comparable] struct {
	want T
	fh   FormatHelper[T]
}

func (nem *NotEqualMatcher[T]) Match(got T) bool {
	return got != nem.want
}

func (nem *NotEqualMatcher[T]) Explain(matched bool, got T) string {
	expected := fmt.Sprintf("got != %s", nem.fh.Format(nem.want))
	if matched {
		return expected
	}
	return fmt.Sprintf("%s != %s", nem.fh.Format(got), nem.fh.Format(nem.want))
}

func (nem *NotEqualMatcher[T]) WithFormat(f func(t T) string) *NotEqualMatcher[T] {
	nem.fh.Set(f)
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

func (aom allOfMatcher[T]) Explain(matched bool, _ T) string {
	actual := "at least one child did not match"
	expected := "all children match"
	return ActualVsExpected(matched, actual, expected)
}

func (aom allOfMatcher[T]) Unwrap(got T) *ResultTree {
	return fanOutMatcherChildren(got, aom.matchers)
}

func fanOutMatcherChildren[T any](got T, matchers []Matcher[T]) *ResultTree {
	if len(matchers) == 0 {
		return nil
	}
	rt := &ResultTree{
		Branches: make([]ResultBranch, len(matchers)),
	}
	for i, matcher := range matchers {
		rt.Branches[i] = ResultBranch{
			Name: fmt.Sprintf("index %d", i),
			ResultTree: ResultTree{
				Root: []Result{MatchResult(got, matcher)},
			},
		}
	}
	return rt
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

func (aom anyOfMatcher[T]) Explain(matched bool, _ T) string {
	actual := "no child matches"
	expected := "at least one child matches"
	return ActualVsExpected(matched, actual, expected)
}

func (aom anyOfMatcher[T]) Unwrap(got T) *ResultTree {
	return fanOutMatcherChildren(got, aom.matchers)
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

func (dm derefMatcher[T]) Explain(matched bool, got *T) string {
	var actual string
	if got == nil {
		actual = "got = nil"
	} else {
		actual = fmt.Sprintf("*got = %v", *got)
	}
	expected := "*got matches child"
	return ActualVsExpected(matched, actual, expected)
}

func (dm derefMatcher[T]) Unwrap(got *T) *ResultTree {
	if got == nil {
		return nil
	}
	return NewResultTreeRoot(MatchResult(*got, dm.matcher))
}

func PointerEqual[T comparable](want *T) *PointerEqualMatcher[T] {
	return &PointerEqualMatcher[T]{want: want}
}

type PointerEqualMatcher[T comparable] struct {
	want *T
	fh   FormatHelper[T]
}

func (pem *PointerEqualMatcher[T]) Match(got *T) bool {
	if pem.want == nil || got == nil {
		return pem.want == got
	}
	return *pem.want == *got
}

func (pem *PointerEqualMatcher[T]) Explain(matched bool, got *T) string {
	var expected string
	if pem.want == nil {
		expected = "got == nil"
	} else {
		expected = fmt.Sprintf("*got == %v", pem.fh.Format(*pem.want))
	}
	var actual string
	if got == nil {
		actual = "got == nil"
	} else {
		actual = fmt.Sprintf("*got = %v", pem.fh.Format(*got))
	}
	return ActualVsExpected(matched, actual, expected)
}

func (pem *PointerEqualMatcher[T]) WithFormat(f func(t T) string) *PointerEqualMatcher[T] {
	pem.fh.Set(f)
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

// TODO: support Unwrap() on ElementsMatcher.

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

func (em *ElementsMatcher[T]) Explain(matched bool, _ []T) string {
	var actual string
	if matched {
		actual = "all elements matched"
	} else {
		actual = "at least one element did not match"
	}
	expected := "all elements match"
	return ActualVsExpected(matched, actual, expected)
}

func Slice[T any](want []T, factory func(t T) Matcher[T]) Matcher[[]T] {
	matchers := make([]Matcher[T], len(want))
	for i, w := range want {
		matchers[i] = factory(w)
	}
	return Elements(matchers...)
}

func AsAny[T any](matcher Matcher[T]) Matcher[any] {
	return anyMatcher[T]{matcher: matcher}
}

type anyMatcher[T any] struct {
	matcher Matcher[T]
}

func (am anyMatcher[T]) Match(got any) bool {
	t, ok := got.(T)
	if !ok {
		return false
	}
	return am.matcher.Match(t)
}

func TypeName[T any]() string {
	return reflect.TypeFor[T]().Name()
}

func (am anyMatcher[T]) Explain(matched bool, got any) string {
	var actual string
	expected := fmt.Sprintf("got is of type %s and matches child", TypeName[T]())
	if matched {
		actual = expected
	} else {
		if _, ok := got.(T); !ok {
			actual = fmt.Sprintf("got is of type %s", reflect.TypeOf(got).String())
		} else {
			actual = fmt.Sprintf("got is of type %s but does not match child", TypeName[T]())
		}
	}
	return ActualVsExpected(matched, actual, expected)
}

func (am anyMatcher[T]) Unwrap(got any) *ResultTree {
	t, ok := got.(T)
	if !ok {
		return nil
	}
	return NewResultTreeRoot(MatchResult(t, am.matcher))
}

func AsType[T any](matcher Matcher[any]) Matcher[T] {
	return typeMatcher[T]{matcher: matcher}
}

type typeMatcher[T any] struct {
	matcher Matcher[any]
}

func (tm typeMatcher[T]) Match(got T) bool {
	return tm.matcher.Match(got)
}

func (tm typeMatcher[T]) Explain(matched bool, _ T) string {
	expected := "got matches child"
	var actual string
	if matched {
		actual = expected
	} else {
		actual = fmt.Sprintf("got of type %s does not match child", TypeName[T]())
	}
	return ActualVsExpected(matched, actual, expected)
}

func (tm typeMatcher[T]) Unwrap(got T) *ResultTree {
	return NewResultTreeRoot(MatchResult(any(got), tm.matcher))
}
