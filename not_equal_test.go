package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
	"github.com/krelinga/go-typemap"
)

func TestNotEqualTm(t *testing.T) {
	goldie := newGoldie(t)
	tm := struct {
		typemap.StringFunc[int]
		typemap.DefaultCompare[int]
	}{
		StringFunc: match.DefaultString[int],
	}
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "not_equal_values",
			matcher: match.NotEqualTm(tm, 42),
			value:   43,
			want:    true,
		},
		{
			name:    "equal_values",
			matcher: match.NotEqualTm(tm, 42),
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
