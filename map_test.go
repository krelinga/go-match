package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)

func TestMapHas(t *testing.T) {
	m := match.NewMapHas(10, match.Equal[string]{X: "value"})
	if !match.Match(map[int]string{10: "value"}, m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.Explain(map[int]string{10: "value"}, m))
	t.Logf("\n%s", match.Explain(map[int]string{10: "other"}, m))
	t.Logf("\n%s", match.Explain(map[int]string{11: "value"}, m))
	t.Logf("\n%s", match.Explain(map[int]string{}, m))
}
