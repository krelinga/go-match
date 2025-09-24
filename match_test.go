package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)

func Test(t *testing.T) {
	t.Run("Single Equals", func(t *testing.T) {
		result := match.Match(1, match.Equal(2))
		if result.Match {
			t.Errorf("Expected no match, got match")
		}
		t.Logf("\n%s\n", result)
	})

	t.Run("AllOf with one failing", func(t *testing.T) {
		result := match.Match(1, match.AllOf(
			match.Equal(1),
			match.Equal(2),
		))
		if result.Match {
			t.Errorf("Expected no match, got match")
		}
		t.Logf("\n%s\n", result)
	})

	t.Run("AnyOf with one passing", func(t *testing.T) {
		result := match.Match(1, match.AnyOf(
			match.Equal(1),
			match.Equal(2),
		))
		if !result.Match {
			t.Errorf("Expected match, got no match")
		}
		t.Logf("\n%s\n", result)
	})
}
