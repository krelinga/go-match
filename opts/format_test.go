package opts_test

import (
	"reflect"
	"testing"

	"github.com/krelinga/go-match/opts"
)

func TestFmt(t *testing.T) {
	o := opts.NewOptions()
	o = opts.SetFmtFunc(nil, func(v reflect.Value) (string, error) {
		return "formatted", nil
	})(o)
	gotStr, gotErr := opts.Fmt(o, reflect.ValueOf(42))
	if gotErr != nil {
		t.Errorf("Fmt() error = %v, wantErr %v", gotErr, false)
	}
	if gotStr != "formatted" {
		t.Errorf("Fmt() = %v, want %v", gotStr, "formatted")
	}
}
