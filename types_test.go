package match

import (
	"reflect"
	"testing"
)

func assertImplements[Got any, Want any](t *testing.T) {
	t.Helper()
	var gotZero Got
	if _, ok := any(gotZero).(Want); !ok {
		got := reflect.TypeFor[Got]().String()
		want := reflect.TypeFor[Want]().String()
		t.Errorf("%s does not implement %s", got, want)
	}
}

func TestEqual(t *testing.T) {
	assertImplements[Equal[int], Matcher[int]](t)
	assertImplements[Equal[int], Explainer[int]](t)
}

func TestNotEqual(t *testing.T) {
	assertImplements[NotEqual[int], Matcher[int]](t)
	assertImplements[NotEqual[int], Explainer[int]](t)
}

func TestAllOf(t *testing.T) {
	assertImplements[AllOf[int], Matcher[int]](t)
	assertImplements[AllOf[int], Explainer[int]](t)
}

func TestAnyOf(t *testing.T) {
	assertImplements[AnyOf[int], Matcher[int]](t)
	assertImplements[AnyOf[int], Explainer[int]](t)
}

func TestWhenDeref(t *testing.T) {
	assertImplements[WhenDeref[int], Matcher[*int]](t)
	assertImplements[WhenDeref[int], Explainer[*int]](t)
}

func TestSliceElems(t *testing.T) {
	assertImplements[SliceElems[int], Matcher[[]int]](t)
	assertImplements[SliceElems[int], Explainer[[]int]](t)
}
