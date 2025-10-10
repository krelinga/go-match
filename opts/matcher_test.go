package opts_test

import (
	"reflect"
	"testing"

	"github.com/krelinga/go-match/opts"
)

func match(t *testing.T, vs opts.Values, m opts.Func) bool {
	t.Helper()
	opts := opts.NewOptions()
	matched, err := m(opts, vs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !matched {
		t.Errorf("expected match, got no match")
	}
	return matched
}

func TestEqual(t *testing.T) {
	input := opts.NewValues1(42)
	ff := func(v reflect.Value) (string, error) {
		return "formatted", nil
	}
	match(t, input, opts.WithOptions(
		[]opts.Option{opts.WithFmtFunc(nil, ff)},
		opts.Equal(41)))
}
