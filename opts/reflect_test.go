package opts_test

import (
	"reflect"
	"testing"

	"github.com/krelinga/go-match/opts"
)

type TestInter interface {
	Get() string
}

type TestStr string

func (ts TestStr) Get() string {
	return string(ts)
}

func TestValueOf(t *testing.T) {
	var testInter TestInter = TestStr("test")
	val := opts.ValueOf(testInter)
	if val.Interface().(TestInter).Get() != "test" {
		t.Errorf("ValueOf failed")
	}
	if val.Type() != reflect.TypeFor[TestInter]() {
		t.Errorf("ValueOf type mismatch")
	}
}