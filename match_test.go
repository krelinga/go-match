package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)

func TestEqual(t *testing.T) {
	m := match.Equal(int(42))
	if !match.Match(int(42), m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.MatchResult(int(42), m))
	t.Logf("\n%s", match.MatchResult(int(43), m))
}

func TestAllOf(t *testing.T) {
	m := match.AllOf(
		match.Equal(int(42)),
		match.NotEqual(int(43)),
	)
	if !match.Match(int(42), m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.MatchResult(int(42), m))
	t.Logf("\n%s", match.MatchResult(int(41), m))
}
