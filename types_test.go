package match

import (
	"reflect"
	"testing"
)

func implements[Got any, Want any](t *testing.T) {
	t.Helper()
	var gotZero Got
	if _, ok := any(gotZero).(Want); !ok {
		got := reflect.TypeFor[Got]().String()
		want := reflect.TypeFor[Want]().String()
		t.Errorf("%s does not implement %s", got, want)
	}
}

func TestEqual(t *testing.T) {
	implements[*EqualMatcher[int], Matcher[int]](t)
	implements[*EqualMatcher[int], Explainer[int]](t)
}

func TestNotEqual(t *testing.T) {
	implements[*NotEqualMatcher[int], Matcher[int]](t)
	implements[*NotEqualMatcher[int], Explainer[int]](t)
}

func TestAllOf(t *testing.T) {
	implements[*allOfMatcher[int], Matcher[int]](t)
	implements[*allOfMatcher[int], Explainer[int]](t)
	implements[*allOfMatcher[int], Unwrapper[int]](t)
}

func TestAnyOf(t *testing.T) {
	implements[*anyOfMatcher[int], Matcher[int]](t)
	implements[*anyOfMatcher[int], Explainer[int]](t)
	implements[*anyOfMatcher[int], Unwrapper[int]](t)
}

func TestDeref(t *testing.T) {
	implements[derefMatcher[int], Matcher[*int]](t)
	implements[derefMatcher[int], Explainer[*int]](t)
	implements[derefMatcher[int], Unwrapper[*int]](t)
}

func TestPointerEqual(t *testing.T) {
	implements[*PointerEqualMatcher[int], Matcher[*int]](t)
	implements[*PointerEqualMatcher[int], Explainer[*int]](t)
}

func TestElements(t *testing.T) {
	implements[*ElementsMatcher[int], Matcher[[]int]](t)
	implements[*ElementsMatcher[int], Explainer[[]int]](t)
	// TODO: Uncomment when ElementsMatcher implements Unwrapper.
	// implements[*ElementsMatcher[int], Unwrapper[[]int]](t)
}

func TestAsAny(t *testing.T) {
	implements[anyMatcher[int], Matcher[any]](t)
	implements[anyMatcher[int], Explainer[any]](t)
	implements[anyMatcher[int], Unwrapper[any]](t)
}

func TestAsType(t *testing.T) {
	implements[typeMatcher[int], Matcher[int]](t)
	implements[typeMatcher[int], Explainer[int]](t)
	implements[typeMatcher[int], Unwrapper[int]](t)
}
