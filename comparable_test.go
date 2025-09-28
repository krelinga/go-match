package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)

func TestEqual(t *testing.T) {
	m := match.Equal[int]{X: 42}
	if !match.Match(42, m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.Explain(42, m))
	t.Logf("\n%s", match.Explain(43, m))
}
