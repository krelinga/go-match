package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
	"github.com/krelinga/go-typemap"
)

func TestGreaterThanTm(t *testing.T) {
	goldie := newGoldie(t)
	tm := struct {
		typemap.StringFunc[int]
		typemap.DefaultOrder[int]
	}{
		StringFunc: match.DefaultString[int](),
	}
	tests := []struct {
		name    string
		matcher match.Matcher[int]
		value   int
		want    bool
	}{
		{
			name:    "less_than",
			matcher: match.GreaterThanTm(tm, 42),
			value:   41,
			want:    false,
		},
		{
			name:    "equal",
			matcher: match.GreaterThanTm(tm, 42),
			value:   42,
			want:    false,
		},
		{
			name:    "greater_than",
			matcher: match.GreaterThanTm(tm, 42),
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
