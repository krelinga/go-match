package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)

func TestSliceElems(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.SliceElems[int], match.Matcher[[]int]](t)
		assertImplements[match.SliceElems[int], match.Explainer[[]int]](t)
	})
	
	t.Run("Ordered", func(t *testing.T) {
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
	})

	t.Run("Unordered", func(t *testing.T) {
		m := match.SliceElems[int]{
			M: []match.Matcher[int]{
				match.Equal[int]{X: 1},
				match.Equal[int]{X: 2},
				match.Equal[int]{X: 3},
			},
			InAnyOrder: true,
		}
		if !match.Match([]int{3, 1, 2}, m) {
			t.Error("Expected match")
		}
		t.Logf("\n%s", match.Explain([]int{3, 1, 2}, m))
		t.Logf("\n%s", match.Explain([]int{1, 2, 4}, m))
		t.Logf("\n%s", match.Explain([]int{1, 2, 3, 4}, m))
		t.Logf("\n%s", match.Explain([]int{1, 2}, m))
	})
}

func TestSliceLen(t *testing.T) {
	m := match.SliceLen[string]{
		M: match.Equal[int]{X: 3},
	}
	if !match.Match([]string{"a", "b", "c"}, m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.Explain([]string{"a", "b", "c"}, m))
	t.Logf("\n%s", match.Explain([]string{"a", "b"}, m))
	t.Logf("\n%s", match.Explain([]string{"a", "b", "c", "d"}, m))
}

func TestSliceNil(t *testing.T) {
	m := match.SliceNil[string]{}
	if !match.Match(nil, m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.Explain(nil, m))
	t.Logf("\n%s", match.Explain([]string{}, m))
}

func TestSliceHas(t *testing.T) {
	m := match.SliceHas[int]{
		M: match.Equal[int]{X: 42},
	}
	if !match.Match([]int{1, 42, 3}, m) {
		t.Error("Expected match")
	}
	t.Logf("\n%s", match.Explain([]int{1, 42, 3}, m))
	t.Logf("\n%s", match.Explain([]int{1, 2, 3}, m))
}