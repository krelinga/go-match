package match_test

import (
	"fmt"
	"testing"

	"github.com/krelinga/go-match"
)

func TestLessThan(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.LessThan[int], match.Matcher[int]](t)
		assertImplements[match.LessThan[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.LessThan[int]
		val  int
		want bool
	}{
		{
			name: "match",
			m:    match.LessThan[int]{X: 10},
			val:  5,
			want: true,
		},
		{
			name: "no_match_equal",
			m:    match.LessThan[int]{X: 10},
			val:  10,
			want: false,
		},
		{
			name: "no_match_greater",
			m:    match.LessThan[int]{X: 10},
			val:  15,
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewLessThan(10),
			val:  5,
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewLessThan(10),
			val:  15,
			want: false,
		},
		{
			name: "match_format",
			m:    match.LessThan[int]{X: 10, Format: func(i int) string { return fmt.Sprintf("(%d)", i) }},
			val:  5,
			want: true,
		},
		{
			name: "no_match_format",
			m:    match.LessThan[int]{X: 10, Format: func(i int) string { return fmt.Sprintf("(%d)", i) }},
			val:  15,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := match.Match(tt.val, tt.m)
			if got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
			goldie.Assert(t, tt.name, []byte(match.Explain(tt.val, tt.m)))
		})
	}
}

func TestLessThanOrEqual(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.LessThanOrEqual[int], match.Matcher[int]](t)
		assertImplements[match.LessThanOrEqual[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.LessThanOrEqual[int]
		val  int
		want bool
	}{
		{
			name: "match_less",
			m:    match.LessThanOrEqual[int]{X: 10},
			val:  5,
			want: true,
		},
		{
			name: "match_equal",
			m:    match.LessThanOrEqual[int]{X: 10},
			val:  10,
			want: true,
		},
		{
			name: "no_match",
			m:    match.LessThanOrEqual[int]{X: 10},
			val:  15,
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewLessThanOrEqual(10),
			val:  5,
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewLessThanOrEqual(10),
			val:  15,
			want: false,
		},
		{
			name: "match_format",
			m:    match.LessThanOrEqual[int]{X: 10, Format: func(i int) string { return fmt.Sprintf("(%d)", i) }},
			val:  10,
			want: true,
		},
		{
			name: "no_match_format",
			m:    match.LessThanOrEqual[int]{X: 10, Format: func(i int) string { return fmt.Sprintf("(%d)", i) }},
			val:  15,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := match.Match(tt.val, tt.m)
			if got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
			goldie.Assert(t, tt.name, []byte(match.Explain(tt.val, tt.m)))
		})
	}
}

func TestGreaterThan(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.GreaterThan[int], match.Matcher[int]](t)
		assertImplements[match.GreaterThan[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.GreaterThan[int]
		val  int
		want bool
	}{
		{
			name: "match",
			m:    match.GreaterThan[int]{X: 10},
			val:  15,
			want: true,
		},
		{
			name: "no_match_equal",
			m:    match.GreaterThan[int]{X: 10},
			val:  10,
			want: false,
		},
		{
			name: "no_match_less",
			m:    match.GreaterThan[int]{X: 10},
			val:  5,
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewGreaterThan(10),
			val:  15,
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewGreaterThan(10),
			val:  5,
			want: false,
		},
		{
			name: "match_format",
			m:    match.GreaterThan[int]{X: 10, Format: func(i int) string { return fmt.Sprintf("(%d)", i) }},
			val:  15,
			want: true,
		},
		{
			name: "no_match_format",
			m:    match.GreaterThan[int]{X: 10, Format: func(i int) string { return fmt.Sprintf("(%d)", i) }},
			val:  5,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := match.Match(tt.val, tt.m)
			if got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
			goldie.Assert(t, tt.name, []byte(match.Explain(tt.val, tt.m)))
		})
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.GreaterThanOrEqual[int], match.Matcher[int]](t)
		assertImplements[match.GreaterThanOrEqual[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.GreaterThanOrEqual[int]
		val  int
		want bool
	}{
		{
			name: "match_greater",
			m:    match.GreaterThanOrEqual[int]{X: 10},
			val:  15,
			want: true,
		},
		{
			name: "match_equal",
			m:    match.GreaterThanOrEqual[int]{X: 10},
			val:  10,
			want: true,
		},
		{
			name: "no_match",
			m:    match.GreaterThanOrEqual[int]{X: 10},
			val:  5,
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewGreaterThanOrEqual(10),
			val:  15,
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewGreaterThanOrEqual(10),
			val:  5,
			want: false,
		},
		{
			name: "match_format",
			m:    match.GreaterThanOrEqual[int]{X: 10, Format: func(i int) string { return fmt.Sprintf("(%d)", i) }},
			val:  10,
			want: true,
		},
		{
			name: "no_match_format",
			m:    match.GreaterThanOrEqual[int]{X: 10, Format: func(i int) string { return fmt.Sprintf("(%d)", i) }},
			val:  5,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := match.Match(tt.val, tt.m)
			if got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
			goldie.Assert(t, tt.name, []byte(match.Explain(tt.val, tt.m)))
		})
	}
}