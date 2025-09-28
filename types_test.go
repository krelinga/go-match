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
	implements[Equal[int], Matcher[int]](t)
	implements[Equal[int], Explainer[int]](t)
}

func TestNotEqual(t *testing.T) {
	implements[NotEqual[int], Matcher[int]](t)
	implements[NotEqual[int], Explainer[int]](t)
}

func TestAllOf(t *testing.T) {
	implements[AllOf[int], Matcher[int]](t)
	implements[AllOf[int], Explainer[int]](t)
}

func TestAnyOf(t *testing.T) {
	implements[AnyOf[int], Matcher[int]](t)
	implements[AnyOf[int], Explainer[int]](t)
}

func TestWhenDeref(t *testing.T) {
	implements[WhenDeref[int], Matcher[*int]](t)
	implements[WhenDeref[int], Explainer[*int]](t)
}

func TestSliceElems(t *testing.T) {
	implements[SliceElems[int], Matcher[[]int]](t)
	implements[SliceElems[int], Explainer[[]int]](t)
}
