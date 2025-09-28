package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)


func TestSliceElems(t *testing.T) {
	m := match.NewSliceElems(
		match.Equal[int]{X: 1},
		match.Equal[int]{X: 2},
		match.Equal[int]{X: 3},
	)
	if !match.Match([]int{1, 2, 3}, m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.Explain([]int{1, 2, 3}, m))
	t.Logf("\n%s", match.Explain([]int{1, 2, 4}, m))
	t.Logf("\n%s", match.Explain([]int{1, 2, 3, 4}, m))
	t.Logf("\n%s", match.Explain([]int{1, 2}, m))
}