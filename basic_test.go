package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)

func TestEqual(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "equal_values",
			matcher: match.Equal(42),
			value:   42,
			want:    true,
		},
		{
			name:    "not_equal_values",
			matcher: match.Equal(42),
			value:   43,
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

func TestNotEqual(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "not_equal_values",
			matcher: match.NotEqual(42),
			value:   43,
			want:    true,
		},
		{
			name:    "equal_values",
			matcher: match.NotEqual(42),
			value:   42,
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

func TestLessThan(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "less_than",
			matcher: match.LessThan(42),
			value:   41,
			want:    true,
		},
		{
			name:    "equal",
			matcher: match.LessThan(42),
			value:   42,
			want:    false,
		},
		{
			name:    "greater_than",
			matcher: match.LessThan(42),
			value:   43,
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

func TestLessThanOrEqual(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "less_than",
			matcher: match.LessThanOrEqual(42),
			value:   41,
			want:    true,
		},
		{
			name:    "equal",
			matcher: match.LessThanOrEqual(42),
			value:   42,
			want:    true,
		},
		{
			name:    "greater_than",
			matcher: match.LessThanOrEqual(42),
			value:   43,
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

func TestGreaterThan(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "less_than",
			matcher: match.GreaterThan(42),
			value:   41,
			want:    false,
		},
		{
			name:    "equal",
			matcher: match.GreaterThan(42),
			value:   42,
			want:    false,
		},
		{
			name:    "greater_than",
			matcher: match.GreaterThan(42),
			value:   43,
			want:    true,
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

func TestGreaterThanOrEqual(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "less_than",
			matcher: match.GreaterThanOrEqual(42),
			value:   41,
			want:    false,
		},
		{
			name:    "equal",
			matcher: match.GreaterThanOrEqual(42),
			value:   42,
			want:    true,
		},
		{
			name:    "greater_than",
			matcher: match.GreaterThanOrEqual(42),
			value:   43,
			want:    true,
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
