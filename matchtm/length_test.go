package matchtm_test

import (
	"testing"

	"github.com/krelinga/go-match"
	"github.com/krelinga/go-match/matchtm"
	"github.com/krelinga/go-typemap"
)

func TestLength(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name string
		run  func() (bool, string)
		want bool
	}{
		{
			name: "slice_matched",
			run: func() (bool, string) {
				s := []string{"a", "b", "c"}
				m := matchtm.Length(typemap.ForSlice[string]{}, match.GreaterThan(2))
				return m.Match(s)
			},
			want: true,
		},
		{
			name: "slice_not_matched",
			run: func() (bool, string) {
				s := []string{"a", "b", "c"}
				m := matchtm.Length(typemap.ForSlice[string]{}, match.LessThan(2))
				return m.Match(s)
			},
			want: false,
		},
		{
			name: "map_matched",
			run: func() (bool, string) {
				m := map[string]int{"a": 1, "b": 2, "c": 3}
				matcher := matchtm.Length(typemap.ForMap[string, int]{}, match.GreaterThan(2))
				return matcher.Match(m)
			},
			want: true,
		},
		{
			name: "map_not_matched",
			run: func() (bool, string) {
				m := map[string]int{"a": 1, "b": 2, "c": 3}
				matcher := matchtm.Length(typemap.ForMap[string, int]{}, match.LessThan(2))
				return matcher.Match(m)
			},
			want: false,
		},
		{
			name: "string_matched",
			run: func() (bool, string) {
				s := "hello"
				m := matchtm.Length(typemap.ForString{}, match.GreaterThan(3))
				return m.Match(s)
			},
			want: true,
		},
		{
			name: "string_not_matched",
			run: func() (bool, string) {
				s := "hello"
				m := matchtm.Length(typemap.ForString{}, match.LessThan(3))
				return m.Match(s)
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, explanation := tt.run()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}
