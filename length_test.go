package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
	"github.com/krelinga/go-typemap"
)

func TestLengthTm(t *testing.T) {
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
			name:    "length_equal",
			matcher: match.LengthTm(tm, match.Equal(3)),
			value:   []int{1, 2, 3},
			want:    true,
		},
		{
			name:    "length_not_equal",
			matcher: match.LengthTm(tm, match.Equal(5)),
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

func TestStringLikeLength(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[string]
		value   string
		want    bool
	}{
		{
			name:    "length_equal",
			matcher: match.StringLikeLength[string](match.Equal(5)),
			value:   "hello",
			want:    true,
		},
		{
			name:    "length_not_equal",
			matcher: match.StringLikeLength[string](match.Equal(3)),
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

func TestStringLength(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[string]
		value   string
		want    bool
	}{
		{
			name:    "length_equal",
			matcher: match.StringLength(match.Equal(5)),
			value:   "hello",
			want:    true,
		},
		{
			name:    "length_not_equal",
			matcher: match.StringLength(match.Equal(3)),
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

func TestSliceLength(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[[]int]
		value   []int
		want    bool
	}{
		{
			name:    "length_equal",
			matcher: match.SliceLength[[]int](match.Equal(3)),
			value:   []int{1, 2, 3},
			want:    true,
		},
		{
			name:    "length_not_equal",
			matcher: match.SliceLength[[]int](match.Equal(5)),
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

func TestMapLength(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[map[string]int]
		value   map[string]int
		want    bool
	}{
		{
			name:    "length_equal",
			matcher: match.MapLength[map[string]int](match.Equal(2)),
			value:   map[string]int{"foo": 1, "bar": 2},
			want:    true,
		},
		{
			name:    "length_not_equal",
			matcher: match.MapLength[map[string]int](match.Equal(5)),
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
