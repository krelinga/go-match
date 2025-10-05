package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
	"github.com/krelinga/go-typemap"
)

func TestStringContainsTm(t *testing.T) {
	type CustomString string

	tm := typemap.StringFunc[CustomString](func(s CustomString) string {
		return string(s)
	})

	tests := []struct {
		name      string
		input     CustomString
		substr    string
		wantMatch bool
	}{
		{
			name:      "contains_substring",
			input:     "hello world",
			substr:    "world",
			wantMatch: true,
		},
		{
			name:      "does_not_contain_substring",
			input:     "hello world",
			substr:    "foo",
			wantMatch: false,
		},
		{
			name:      "empty_substring_matches_any_string",
			input:     "hello",
			substr:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_non_empty_substring",
			input:     "",
			substr:    "test",
			wantMatch: false,
		},
		{
			name:      "empty_string_with_empty_substring",
			input:     "",
			substr:    "",
			wantMatch: true,
		},
		{
			name:      "case_sensitive_match",
			input:     "Hello World",
			substr:    "hello",
			wantMatch: false,
		},
		{
			name:      "exact_match",
			input:     "test",
			substr:    "test",
			wantMatch: true,
		},
		{
			name:      "substring_at_beginning",
			input:     "testing123",
			substr:    "test",
			wantMatch: true,
		},
		{
			name:      "substring_at_end",
			input:     "123test",
			substr:    "test",
			wantMatch: true,
		},
	}

	goldie := newGoldie(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := match.StringContainsTm(tm, tt.substr)
			matched, explanation := matcher.Match(tt.input)

			if matched != tt.wantMatch {
				t.Errorf("StringContainsTm().Match() matched = %v, want %v", matched, tt.wantMatch)
			}

			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}

func TestStringLikeContains(t *testing.T) {
	type MyString string

	tests := []struct {
		name      string
		input     MyString
		substr    string
		wantMatch bool
	}{
		{
			name:      "contains_substring",
			input:     "hello world",
			substr:    "world",
			wantMatch: true,
		},
		{
			name:      "does_not_contain_substring",
			input:     "hello world",
			substr:    "foo",
			wantMatch: false,
		},
		{
			name:      "empty_substring",
			input:     "test",
			substr:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_substring",
			input:     "",
			substr:    "test",
			wantMatch: false,
		},
		{
			name:      "both_empty",
			input:     "",
			substr:    "",
			wantMatch: true,
		},
		{
			name:      "special_characters",
			input:     "hello@world.com",
			substr:    "@world",
			wantMatch: true,
		},
		{
			name:      "unicode_characters",
			input:     "こんにちは世界",
			substr:    "世界",
			wantMatch: true,
		},
		{
			name:      "partial_unicode_match",
			input:     "こんにちは世界",
			substr:    "hello",
			wantMatch: false,
		},
	}

	goldie := newGoldie(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := match.StringLikeContains[MyString](tt.substr)
			matched, explanation := matcher.Match(tt.input)

			if matched != tt.wantMatch {
				t.Errorf("StringLikeContains().Match() matched = %v, want %v", matched, tt.wantMatch)
			}

			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}

func TestStringContains(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		substr    string
		wantMatch bool
	}{
		{
			name:      "simple_contains",
			input:     "hello world",
			substr:    "world",
			wantMatch: true,
		},
		{
			name:      "does_not_contain",
			input:     "hello world",
			substr:    "universe",
			wantMatch: false,
		},
		{
			name:      "empty_substring_always_matches",
			input:     "anything",
			substr:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_empty_substring",
			input:     "",
			substr:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_non_empty_substring",
			input:     "",
			substr:    "test",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_different_case",
			input:     "Hello World",
			substr:    "WORLD",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_same_case",
			input:     "Hello World",
			substr:    "World",
			wantMatch: true,
		},
		{
			name:      "substring_longer_than_input",
			input:     "hi",
			substr:    "hello",
			wantMatch: false,
		},
		{
			name:      "exact_match",
			input:     "exact",
			substr:    "exact",
			wantMatch: true,
		},
		{
			name:      "whitespace_handling",
			input:     "hello world",
			substr:    " ",
			wantMatch: true,
		},
		{
			name:      "newline_characters",
			input:     "line1\nline2",
			substr:    "\n",
			wantMatch: true,
		},
		{
			name:      "tab_characters",
			input:     "col1\tcol2",
			substr:    "\t",
			wantMatch: true,
		},
		{
			name:      "multiple_occurrences",
			input:     "test test test",
			substr:    "test",
			wantMatch: true,
		},
		{
			name:      "overlapping_pattern",
			input:     "aaaa",
			substr:    "aa",
			wantMatch: true,
		},
	}

	goldie := newGoldie(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := match.StringContains(tt.substr)
			matched, explanation := matcher.Match(tt.input)

			if matched != tt.wantMatch {
				t.Errorf("StringContains().Match() matched = %v, want %v", matched, tt.wantMatch)
			}

			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}
