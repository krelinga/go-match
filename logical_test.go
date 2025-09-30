package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
)

func TestAllOf(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name: "all_matchers_match",
			matcher: match.AllOf(
				match.LessThan(100),
				match.GreaterThan(10),
				match.NotEqual(50),
			),
			value: 42,
			want:  true,
		},
		{
			name: "one_matcher_does_not_match",
			matcher: match.AllOf(
				match.LessThan(100),
				match.GreaterThan(10),
				match.NotEqual(42),
			),
			value: 42,
			want:  false,
		},
		{
			name: "multiple_matchers_do_not_match",
			matcher: match.AllOf(
				match.LessThan(100),
				match.GreaterThan(50),
				match.NotEqual(42),
			),
			value: 42,
			want:  false,
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

func TestAnyOf(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name: "all_matchers_match",
			matcher: match.AnyOf(
				match.LessThan(100),
				match.GreaterThan(10),
				match.NotEqual(50),
			),
			value: 42,
			want:  true,
		},
		{
			name: "one_matcher_matches",
			matcher: match.AnyOf(
				match.LessThan(10),
				match.GreaterThan(100),
				match.NotEqual(50),
			),
			value: 42,
			want:  true,
		},
		{
			name: "no_matchers_match",
			matcher: match.AnyOf(
				match.LessThan(10),
				match.GreaterThan(100),
				match.Equal(50),
			),
			value: 42,
			want:  false,
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

func TestNot(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "negated_matcher_matches",
			matcher: match.Not(match.LessThan(10)),
			value:   42,
			want:    true,
		},
		{
			name:    "negated_matcher_does_not_match",
			matcher: match.Not(match.GreaterThan(10)),
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

func TestAlway(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "always_matches",
			matcher: match.Alway[int](),
			value:   42,
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

func TestNever(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "never_matches",
			matcher: match.Never[int](),
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
