package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
	"github.com/krelinga/go-typemap"
)

func TestStringLikeStartsWithTm(t *testing.T) {
	type CustomString string

	tm := typemap.StringFunc[CustomString](func(s CustomString) string {
		return string(s)
	})

	tests := []struct {
		name      string
		input     CustomString
		prefix    string
		wantMatch bool
	}{
		{
			name:      "starts_with_prefix",
			input:     "hello world",
			prefix:    "hello",
			wantMatch: true,
		},
		{
			name:      "does_not_start_with_prefix",
			input:     "hello world",
			prefix:    "world",
			wantMatch: false,
		},
		{
			name:      "empty_prefix_matches_any_string",
			input:     "hello",
			prefix:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_non_empty_prefix",
			input:     "",
			prefix:    "test",
			wantMatch: false,
		},
		{
			name:      "empty_string_with_empty_prefix",
			input:     "",
			prefix:    "",
			wantMatch: true,
		},
		{
			name:      "case_sensitive_match",
			input:     "Hello World",
			prefix:    "hello",
			wantMatch: false,
		},
		{
			name:      "exact_match",
			input:     "test",
			prefix:    "test",
			wantMatch: true,
		},
		{
			name:      "prefix_longer_than_string",
			input:     "hi",
			prefix:    "hello",
			wantMatch: false,
		},
		{
			name:      "partial_prefix_match",
			input:     "testing123",
			prefix:    "test",
			wantMatch: true,
		},
		{
			name:      "single_character_prefix",
			input:     "apple",
			prefix:    "a",
			wantMatch: true,
		},
		{
			name:      "single_character_no_match",
			input:     "apple",
			prefix:    "b",
			wantMatch: false,
		},
		{
			name:      "whitespace_prefix",
			input:     " hello",
			prefix:    " ",
			wantMatch: true,
		},
		{
			name:      "special_characters",
			input:     "@hello_world",
			prefix:    "@hello",
			wantMatch: true,
		},
		{
			name:      "unicode_characters",
			input:     "こんにちは世界",
			prefix:    "こんにちは",
			wantMatch: true,
		},
		{
			name:      "unicode_no_match",
			input:     "こんにちは世界",
			prefix:    "世界",
			wantMatch: false,
		},
	}

	goldie := newGoldie(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := match.StringLikeStartsWithTm(tm, tt.prefix)
			matched, explanation := matcher.Match(tt.input)

			if matched != tt.wantMatch {
				t.Errorf("StringLikeStartsWithTm().Match() matched = %v, want %v", matched, tt.wantMatch)
			}

			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}

func TestStringLikeStartsWith(t *testing.T) {
	type MyString string

	tests := []struct {
		name      string
		input     MyString
		prefix    string
		wantMatch bool
	}{
		{
			name:      "starts_with_prefix",
			input:     "hello world",
			prefix:    "hello",
			wantMatch: true,
		},
		{
			name:      "does_not_start_with_prefix",
			input:     "hello world",
			prefix:    "world",
			wantMatch: false,
		},
		{
			name:      "empty_prefix",
			input:     "test",
			prefix:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_prefix",
			input:     "",
			prefix:    "test",
			wantMatch: false,
		},
		{
			name:      "both_empty",
			input:     "",
			prefix:    "",
			wantMatch: true,
		},
		{
			name:      "special_characters",
			input:     "hello@world.com",
			prefix:    "hello@",
			wantMatch: true,
		},
		{
			name:      "unicode_characters",
			input:     "こんにちは世界",
			prefix:    "こんに",
			wantMatch: true,
		},
		{
			name:      "partial_unicode_no_match",
			input:     "こんにちは世界",
			prefix:    "世界",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_uppercase",
			input:     "Hello World",
			prefix:    "HELLO",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_correct_case",
			input:     "Hello World",
			prefix:    "Hello",
			wantMatch: true,
		},
		{
			name:      "numeric_prefix",
			input:     "123abc",
			prefix:    "123",
			wantMatch: true,
		},
		{
			name:      "mixed_alphanumeric",
			input:     "abc123def",
			prefix:    "abc1",
			wantMatch: true,
		},
	}

	goldie := newGoldie(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := match.StringLikeStartsWith[MyString](tt.prefix)
			matched, explanation := matcher.Match(tt.input)

			if matched != tt.wantMatch {
				t.Errorf("StringLikeStartsWith().Match() matched = %v, want %v", matched, tt.wantMatch)
			}

			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}

func TestStringStartsWith(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		prefix    string
		wantMatch bool
	}{
		{
			name:      "simple_starts_with",
			input:     "hello world",
			prefix:    "hello",
			wantMatch: true,
		},
		{
			name:      "does_not_start_with",
			input:     "hello world",
			prefix:    "world",
			wantMatch: false,
		},
		{
			name:      "empty_prefix_always_matches",
			input:     "anything",
			prefix:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_empty_prefix",
			input:     "",
			prefix:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_non_empty_prefix",
			input:     "",
			prefix:    "test",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_different_case",
			input:     "Hello World",
			prefix:    "hello",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_same_case",
			input:     "Hello World",
			prefix:    "Hello",
			wantMatch: true,
		},
		{
			name:      "prefix_longer_than_input",
			input:     "hi",
			prefix:    "hello",
			wantMatch: false,
		},
		{
			name:      "exact_match",
			input:     "exact",
			prefix:    "exact",
			wantMatch: true,
		},
		{
			name:      "whitespace_handling",
			input:     " hello world",
			prefix:    " ",
			wantMatch: true,
		},
		{
			name:      "newline_characters",
			input:     "\nline2",
			prefix:    "\n",
			wantMatch: true,
		},
		{
			name:      "tab_characters",
			input:     "\tcol2",
			prefix:    "\t",
			wantMatch: true,
		},
		{
			name:      "single_character_match",
			input:     "test",
			prefix:    "t",
			wantMatch: true,
		},
		{
			name:      "single_character_no_match",
			input:     "test",
			prefix:    "x",
			wantMatch: false,
		},
		{
			name:      "numeric_string",
			input:     "12345",
			prefix:    "123",
			wantMatch: true,
		},
		{
			name:      "mixed_content",
			input:     "abc123def",
			prefix:    "abc",
			wantMatch: true,
		},
		{
			name:      "repeated_characters",
			input:     "aaabbb",
			prefix:    "aaa",
			wantMatch: true,
		},
		{
			name:      "partial_repeated_characters",
			input:     "aaabbb",
			prefix:    "aa",
			wantMatch: true,
		},
		{
			name:      "url_like_string",
			input:     "https://example.com",
			prefix:    "https://",
			wantMatch: true,
		},
		{
			name:      "path_like_string",
			input:     "/usr/local/bin",
			prefix:    "/usr",
			wantMatch: true,
		},
	}

	goldie := newGoldie(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := match.StringStartsWith(tt.prefix)
			matched, explanation := matcher.Match(tt.input)

			if matched != tt.wantMatch {
				t.Errorf("StringStartsWith().Match() matched = %v, want %v", matched, tt.wantMatch)
			}

			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}
