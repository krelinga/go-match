package match_test

import (
	"fmt"
	"testing"

	"github.com/krelinga/go-match"
)

func TestEqualIntExplain_Golden(t *testing.T) {
	g := newGoldie(t)

	tests := []struct {
		name    string
		matcher match.Equal[int]
		got     int
	}{
		{
			name:    "match_equal_42",
			matcher: match.Equal[int]{X: 42},
			got:     42,
		},
		{
			name:    "no_match_equal_42_got_43",
			matcher: match.Equal[int]{X: 42},
			got:     43,
		},
		{
			name:    "match_equal_negative",
			matcher: match.Equal[int]{X: -10},
			got:     -10,
		},
		{
			name:    "no_match_equal_negative",
			matcher: match.Equal[int]{X: -10},
			got:     5,
		},
		{
			name:    "match_equal_zero",
			matcher: match.Equal[int]{X: 0},
			got:     0,
		},
		{
			name:    "no_match_equal_zero",
			matcher: match.Equal[int]{X: 0},
			got:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := match.Explain(tt.got, tt.matcher)
			g.Assert(t, tt.name, []byte(output))
		})
	}
}

func TestEqualIntWithCustomFormat_Golden(t *testing.T) {
	g := newGoldie(t)

	tests := []struct {
		name    string
		matcher match.Equal[int]
		got     int
	}{
		{
			name: "match_with_hex_format",
			matcher: match.Equal[int]{
				X: 255,
				Format: func(i int) string {
					return fmt.Sprintf("0x%x", i)
				},
			},
			got: 255,
		},
		{
			name: "no_match_with_hex_format",
			matcher: match.Equal[int]{
				X: 255,
				Format: func(i int) string {
					return fmt.Sprintf("0x%x", i)
				},
			},
			got: 256,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := match.Explain(tt.got, tt.matcher)
			g.Assert(t, tt.name, []byte(output))
		})
	}
}
