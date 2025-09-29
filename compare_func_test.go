package match_test

import (
	"cmp"
	"strings"
	"testing"

	"github.com/krelinga/go-match"
)



func TestEqualFunc(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.EqualFunc[int], match.Matcher[int]](t)
		assertImplements[match.EqualFunc[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.EqualFunc[string]
		val  string
		want bool
	}{
		{
			name: "match",
			m:    match.EqualFunc[string]{X: "hello", Func: cmp.Compare[string]},
			val:  "hello",
			want: true,
		},
		{
			name: "no_match",
			m:    match.EqualFunc[string]{X: "hello", Func: cmp.Compare[string]},
			val:  "world",
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewEqualFunc("hello", cmp.Compare[string]),
			val:  "hello",
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewEqualFunc("hello", cmp.Compare[string]),
			val:  "world",
			want: false,
		},
		{
			name: "match_format",
			m:    match.EqualFunc[string]{X: "hello", Func: cmp.Compare[string], Format: strings.ToUpper},
			val:  "hello",
			want: true,
		},
		{
			name: "no_match_format",
			m:    match.EqualFunc[string]{X: "hello", Func: cmp.Compare[string], Format: strings.ToUpper},
			val:  "world",
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

	t.Run("nil_func_panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when Func is nil")
			}
		}()
		m := match.EqualFunc[string]{X: "hello", Func: nil}
		m.Match("hello")
	})
}

func TestNotEqualFunc(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.NotEqualFunc[int], match.Matcher[int]](t)
		assertImplements[match.NotEqualFunc[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.NotEqualFunc[string]
		val  string
		want bool
	}{
		{
			name: "match",
			m:    match.NotEqualFunc[string]{X: "hello", Func: cmp.Compare[string]},
			val:  "world",
			want: true,
		},
		{
			name: "no_match",
			m:    match.NotEqualFunc[string]{X: "hello", Func: cmp.Compare[string]},
			val:  "hello",
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewNotEqualFunc("hello", cmp.Compare[string]),
			val:  "world",
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewNotEqualFunc("hello", cmp.Compare[string]),
			val:  "hello",
			want: false,
		},
		{
			name: "match_format",
			m:    match.NotEqualFunc[string]{X: "hello", Func: cmp.Compare[string], Format: strings.ToUpper},
			val:  "world",
			want: true,
		},
		{
			name: "no_match_format",
			m:    match.NotEqualFunc[string]{X: "hello", Func: cmp.Compare[string], Format: strings.ToUpper},
			val:  "hello",
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

	t.Run("nil_func_panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when Func is nil")
			}
		}()
		m := match.NotEqualFunc[string]{X: "hello", Func: nil}
		m.Match("hello")
	})
}

func TestLessThanFunc(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.LessThanFunc[int], match.Matcher[int]](t)
		assertImplements[match.LessThanFunc[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.LessThanFunc[int]
		val  int
		want bool
	}{
		{
			name: "match",
			m:    match.LessThanFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  5,
			want: true,
		},
		{
			name: "no_match_equal",
			m:    match.LessThanFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  10,
			want: false,
		},
		{
			name: "no_match_greater",
			m:    match.LessThanFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  15,
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewLessThanFunc(10, cmp.Compare[int]),
			val:  5,
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewLessThanFunc(10, cmp.Compare[int]),
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

	t.Run("nil_func_panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when Func is nil")
			}
		}()
		m := match.LessThanFunc[int]{X: 10, Func: nil}
		m.Match(5)
	})
}

func TestLessThanOrEqualFunc(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.LessThanOrEqualFunc[int], match.Matcher[int]](t)
		assertImplements[match.LessThanOrEqualFunc[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.LessThanOrEqualFunc[int]
		val  int
		want bool
	}{
		{
			name: "match_less",
			m:    match.LessThanOrEqualFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  5,
			want: true,
		},
		{
			name: "match_equal",
			m:    match.LessThanOrEqualFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  10,
			want: true,
		},
		{
			name: "no_match",
			m:    match.LessThanOrEqualFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  15,
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewLessThanOrEqualFunc(10, cmp.Compare[int]),
			val:  5,
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewLessThanOrEqualFunc(10, cmp.Compare[int]),
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

	t.Run("nil_func_panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when Func is nil")
			}
		}()
		m := match.LessThanOrEqualFunc[int]{X: 10, Func: nil}
		m.Match(5)
	})
}

func TestGreaterThanFunc(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.GreaterThanFunc[int], match.Matcher[int]](t)
		assertImplements[match.GreaterThanFunc[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.GreaterThanFunc[int]
		val  int
		want bool
	}{
		{
			name: "match",
			m:    match.GreaterThanFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  15,
			want: true,
		},
		{
			name: "no_match_equal",
			m:    match.GreaterThanFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  10,
			want: false,
		},
		{
			name: "no_match_less",
			m:    match.GreaterThanFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  5,
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewGreaterThanFunc(10, cmp.Compare[int]),
			val:  15,
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewGreaterThanFunc(10, cmp.Compare[int]),
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

	t.Run("nil_func_panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when Func is nil")
			}
		}()
		m := match.GreaterThanFunc[int]{X: 10, Func: nil}
		m.Match(15)
	})
}

func TestGreaterThanOrEqualFunc(t *testing.T) {
	t.Run("implements", func(t *testing.T) {
		assertImplements[match.GreaterThanOrEqualFunc[int], match.Matcher[int]](t)
		assertImplements[match.GreaterThanOrEqualFunc[int], match.Explainer[int]](t)
	})

	goldie := newGoldie(t)
	tests := []struct {
		name string
		m    match.GreaterThanOrEqualFunc[int]
		val  int
		want bool
	}{
		{
			name: "match_greater",
			m:    match.GreaterThanOrEqualFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  15,
			want: true,
		},
		{
			name: "match_equal",
			m:    match.GreaterThanOrEqualFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  10,
			want: true,
		},
		{
			name: "no_match",
			m:    match.GreaterThanOrEqualFunc[int]{X: 10, Func: cmp.Compare[int]},
			val:  5,
			want: false,
		},
		{
			name: "match_func",
			m:    match.NewGreaterThanOrEqualFunc(10, cmp.Compare[int]),
			val:  15,
			want: true,
		},
		{
			name: "no_match_func",
			m:    match.NewGreaterThanOrEqualFunc(10, cmp.Compare[int]),
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

	t.Run("nil_func_panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when Func is nil")
			}
		}()
		m := match.GreaterThanOrEqualFunc[int]{X: 10, Func: nil}
		m.Match(15)
	})
}