package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)

func Test(t *testing.T) {
	result := match.Match(1, 2, match.Equals[int]())
	if result.Match {
		t.Errorf("Expected no match, got match")
	}
}