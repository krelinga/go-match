package match_test

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
