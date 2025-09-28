package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)

func TestAllOf(t *testing.T) {
	m := match.NewAllOf(match.Equal[int]{X: 42}, match.NotEqual[int]{X: 43})
	if !match.Match(42, m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.Explain(42, m))
	t.Logf("\n%s", match.Explain(41, m))
}

func TestWhenDeref(t *testing.T) {
	m := match.NewWhenDeref(match.Equal[int]{X: 42})
	val := 42
	if !match.Match(&val, m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.Explain(&val, m))
	otherVal := 43
	t.Logf("\n%s", match.Explain(&otherVal, m))
	t.Logf("\n%s", match.Explain((*int)(nil), m))
}
