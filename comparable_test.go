package match_test

import (
	"strings"
	"testing"

	"github.com/krelinga/go-match"
)

func TestEqual(t *testing.T) {
	goldie := newGoldie(t)
	tests := []struct {
		name    string
		m match.Equal[string]
		val string
		want bool
	}{
		{
			name: "match",
			m:    match.Equal[string]{X: "hello"},
			val:  "hello",
			want: true,
		},
		{
			name: "no_match",
			m:    match.Equal[string]{X: "hello"},
			val:  "world",
			want: false,
		},
		{
			name: "match_func",
			m: match.NewEqual("hello"),
			val:  "hello",
			want: true,
		},
		{
			name: "no_match_func",
			m: match.NewEqual("hello"),
			val:  "world",
			want: false,
		},
		{
			name: "match_format",
			m:    match.Equal[string]{X: "hello", Format: strings.ToUpper},
			val:  "hello",
			want: true,
		},
		{
			name: "no_match_format",
			m:    match.Equal[string]{X: "hello", Format: strings.ToUpper},
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
}
