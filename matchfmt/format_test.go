package matchfmt_test

import (
	"testing"

	"github.com/krelinga/go-match/matchfmt"
)

func TestEmoji(t *testing.T) {
	tests := []struct {
		name     string
		matched  bool
		expected string
	}{
		{
			name:     "matched true returns check mark",
			matched:  true,
			expected: "✅",
		},
		{
			name:     "matched false returns X mark",
			matched:  false,
			expected: "❌",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchfmt.Emoji(tt.matched)
			if result != tt.expected {
				t.Errorf("Emoji(%v) = %q, want %q", tt.matched, result, tt.expected)
			}
		})
	}
}

func TestIndentBy(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		level    int
		expected string
	}{
		{
			name:     "single line with level 0",
			input:    "hello",
			level:    0,
			expected: "hello",
		},
		{
			name:     "single line with level 1",
			input:    "hello",
			level:    1,
			expected: "   hello",
		},
		{
			name:     "single line with level 2",
			input:    "hello",
			level:    2,
			expected: "      hello",
		},
		{
			name:     "multiple lines with level 1",
			input:    "line1\nline2",
			level:    1,
			expected: "   line1\n   line2",
		},
		{
			name:     "multiple lines with level 2",
			input:    "line1\nline2\nline3",
			level:    2,
			expected: "      line1\n      line2\n      line3",
		},
		{
			name:     "empty string",
			input:    "",
			level:    1,
			expected: "   ",
		},
		{
			name:     "string with empty lines",
			input:    "line1\n\nline3",
			level:    1,
			expected: "   line1\n   \n   line3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchfmt.IndentBy(tt.input, tt.level)
			if result != tt.expected {
				t.Errorf("IndentBy(%q, %d) = %q, want %q", tt.input, tt.level, result, tt.expected)
			}
		})
	}
}

func TestIndent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single line",
			input:    "hello",
			expected: "   hello",
		},
		{
			name:     "multiple lines",
			input:    "line1\nline2",
			expected: "   line1\n   line2",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "   ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchfmt.Indent(tt.input)
			if result != tt.expected {
				t.Errorf("Indent(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExplain(t *testing.T) {
	tests := []struct {
		name        string
		matched     bool
		matcherName string
		details     []string
		expected    string
	}{
		{
			name:        "matched true with no details",
			matched:     true,
			matcherName: "EqualTo",
			details:     nil,
			expected:    "✅ EqualTo",
		},
		{
			name:        "matched false with no details",
			matched:     false,
			matcherName: "EqualTo",
			details:     nil,
			expected:    "❌ EqualTo",
		},
		{
			name:        "matched true with single detail",
			matched:     true,
			matcherName: "EqualTo",
			details:     []string{"value matches"},
			expected:    "✅ EqualTo:\n   value matches",
		},
		{
			name:        "matched false with single detail",
			matched:     false,
			matcherName: "EqualTo",
			details:     []string{"value does not match"},
			expected:    "❌ EqualTo:\n   value does not match",
		},
		{
			name:        "matched false with multiple details",
			matched:     false,
			matcherName: "Contains",
			details:     []string{"Expected: substring", "Actual: main string"},
			expected:    "❌ Contains:\n   Expected: substring\n   Actual: main string",
		},
		{
			name:        "empty matcher name",
			matched:     true,
			matcherName: "",
			details:     []string{"some detail"},
			expected:    "✅ :\n   some detail",
		},
		{
			name:        "detail with multiple lines",
			matched:     false,
			matcherName: "ComplexMatcher",
			details:     []string{"line1\nline2"},
			expected:    "❌ ComplexMatcher:\n   line1\n   line2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchfmt.Explain(tt.matched, tt.matcherName, tt.details...)
			if result != tt.expected {
				t.Errorf("Explain(%v, %q, %v) = %q, want %q", tt.matched, tt.matcherName, tt.details, result, tt.expected)
			}
		})
	}
}

func TestActualVsExpected(t *testing.T) {
	tests := []struct {
		name     string
		actual   string
		expected string
		want     string
	}{
		{
			name:     "simple strings",
			actual:   "hello",
			expected: "world",
			want:     "Expected: world\nActual:   hello",
		},
		{
			name:     "empty strings",
			actual:   "",
			expected: "",
			want:     "Expected: \nActual:   ",
		},
		{
			name:     "strings with spaces",
			actual:   "hello world",
			expected: "foo bar",
			want:     "Expected: foo bar\nActual:   hello world",
		},
		{
			name:     "strings with special characters",
			actual:   "hello\nworld",
			expected: "foo\tbar",
			want:     "Expected: foo\tbar\nActual:   hello\nworld",
		},
		{
			name:     "numeric strings",
			actual:   "123",
			expected: "456",
			want:     "Expected: 456\nActual:   123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchfmt.ActualVsExpected(tt.actual, tt.expected)
			if result != tt.want {
				t.Errorf("ActualVsExpected(%q, %q) = %q, want %q", tt.actual, tt.expected, result, tt.want)
			}
		})
	}
}
