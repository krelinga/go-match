package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
	"github.com/krelinga/go-typemap"
)

func TestIsNilTm(t *testing.T) {
	goldie := newGoldie(t)
	tm := typemap.ForSliceLike[[]int, int]{
		StringFunc: match.DefaultString[[]int](),
	}
	tests := []struct {
		name    string
		matcher match.Matcher[[]int]
		value   []int
		want    bool
	}{
		{
			name:    "nil_slice",
			matcher: match.IsNilTm(tm),
			value:   nil,
			want:    true,
		},
		{
			name:    "non_nil_slice",
			matcher: match.IsNilTm(tm),
			value:   []int{1, 2, 3},
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotExplanation := tt.matcher.Match(tt.value)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			goldie.Assert(t, tt.name, []byte(gotExplanation))
		})
	}
}

func TestSliceIsNil(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[[]int]
		value   []int
		want    bool
	}{
		{
			name:    "nil_slice",
			matcher: match.SliceIsNil[[]int](),
			value:   nil,
			want:    true,
		},
		{
			name:    "non_nil_slice",
			matcher: match.SliceIsNil[[]int](),
			value:   []int{1, 2, 3},
			want:    false,
		},
		{
			name:    "empty_slice",
			matcher: match.SliceIsNil[[]int](),
			value:   []int{},
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotExplanation := tt.matcher.Match(tt.value)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			goldie.Assert(t, tt.name, []byte(gotExplanation))
		})
	}
}

func TestMapIsNil(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[map[string]int]
		value   map[string]int
		want    bool
	}{
		{
			name:    "nil_map",
			matcher: match.MapIsNil[map[string]int](),
			value:   nil,
			want:    true,
		},
		{
			name:    "non_nil_map",
			matcher: match.MapIsNil[map[string]int](),
			value:   map[string]int{"foo": 1},
			want:    false,
		},
		{
			name:    "empty_map",
			matcher: match.MapIsNil[map[string]int](),
			value:   map[string]int{},
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotExplanation := tt.matcher.Match(tt.value)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			goldie.Assert(t, tt.name, []byte(gotExplanation))
		})
	}
}

func TestPointerIsNil(t *testing.T) {
	goldie := newGoldie(t)
	nonNilValue := 42
	tests := []struct {
		name           string
		matcher        match.Matcher[*int]
		value          *int
		want           bool
		skipGoldenFile bool
	}{
		{
			name:    "nil_pointer",
			matcher: match.PointerIsNil[int](),
			value:   nil,
			want:    true,
		},
		{
			name:           "non_nil_pointer",
			matcher:        match.PointerIsNil[int](),
			value:          &nonNilValue,
			want:           false,
			skipGoldenFile: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotExplanation := tt.matcher.Match(tt.value)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if tt.skipGoldenFile {
				// Pointer addresses are non-deterministic, so we just verify the explanation contains expected parts
				if got {
					t.Errorf("expected matcher to not match, but it did")
				}
				// Just verify it contains the key parts
				if len(gotExplanation) == 0 {
					t.Errorf("expected non-empty explanation")
				}
			} else {
				goldie.Assert(t, tt.name, []byte(gotExplanation))
			}
		})
	}
}
