package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
	"github.com/krelinga/go-typemap"
)

func TestHasKeyTm(t *testing.T) {
	goldie := newGoldie(t)
	containerTm := typemap.ForMapLike[map[string]int, string, int]{}
	keyTm := struct {
		typemap.StringFunc[string]
	}{
		StringFunc: match.DefaultString[string](),
	}
	tests := []struct {
		name    string
		matcher match.Matcher[map[string]int]
		value   map[string]int
		want    bool
	}{
		{
			name:    "key_exists",
			matcher: match.HasKeyTm(containerTm, keyTm, "foo"),
			value:   map[string]int{"foo": 1, "bar": 2},
			want:    true,
		},
		{
			name:    "key_not_found",
			matcher: match.HasKeyTm(containerTm, keyTm, "baz"),
			value:   map[string]int{"foo": 1, "bar": 2},
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

func TestStringHasIndex(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[string]
		value   string
		want    bool
	}{
		{
			name:    "index_exists",
			matcher: match.StringHasIndex[string](2),
			value:   "hello",
			want:    true,
		},
		{
			name:    "index_out_of_bounds",
			matcher: match.StringHasIndex[string](10),
			value:   "hello",
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

func TestSliceHasIndex(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[[]int]
		value   []int
		want    bool
	}{
		{
			name:    "index_exists",
			matcher: match.SliceHasIndex[[]int](1),
			value:   []int{10, 20, 30},
			want:    true,
		},
		{
			name:    "index_out_of_bounds",
			matcher: match.SliceHasIndex[[]int](5),
			value:   []int{10, 20, 30},
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

func TestMapHasKey(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[map[string]int]
		value   map[string]int
		want    bool
	}{
		{
			name:    "key_exists",
			matcher: match.MapHasKey[map[string]int]("foo"),
			value:   map[string]int{"foo": 1, "bar": 2},
			want:    true,
		},
		{
			name:    "key_not_found",
			matcher: match.MapHasKey[map[string]int]("baz"),
			value:   map[string]int{"foo": 1, "bar": 2},
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
