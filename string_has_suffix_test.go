package match_test

import (
	"testing"

	"github.com/krelinga/go-match"
	"github.com/krelinga/go-typemap"
)

func TestStringLikeHasSuffixTm(t *testing.T) {
	type CustomString string

	tm := typemap.StringFunc[CustomString](func(s CustomString) string {
		return string(s)
	})

	tests := []struct {
		name      string
		input     CustomString
		suffix    string
		wantMatch bool
	}{
		{
			name:      "ends_with_suffix",
			input:     "hello world",
			suffix:    "world",
			wantMatch: true,
		},
		{
			name:      "does_not_end_with_suffix",
			input:     "hello world",
			suffix:    "hello",
			wantMatch: false,
		},
		{
			name:      "empty_suffix_matches_any_string",
			input:     "hello",
			suffix:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_non_empty_suffix",
			input:     "",
			suffix:    "test",
			wantMatch: false,
		},
		{
			name:      "empty_string_with_empty_suffix",
			input:     "",
			suffix:    "",
			wantMatch: true,
		},
		{
			name:      "case_sensitive_match",
			input:     "Hello World",
			suffix:    "world",
			wantMatch: false,
		},
		{
			name:      "exact_match",
			input:     "test",
			suffix:    "test",
			wantMatch: true,
		},
		{
			name:      "suffix_longer_than_string",
			input:     "hi",
			suffix:    "hello",
			wantMatch: false,
		},
		{
			name:      "partial_suffix_match",
			input:     "123testing",
			suffix:    "ting",
			wantMatch: true,
		},
		{
			name:      "single_character_suffix",
			input:     "apple",
			suffix:    "e",
			wantMatch: true,
		},
		{
			name:      "single_character_no_match",
			input:     "apple",
			suffix:    "x",
			wantMatch: false,
		},
		{
			name:      "whitespace_suffix",
			input:     "hello ",
			suffix:    " ",
			wantMatch: true,
		},
		{
			name:      "special_characters",
			input:     "hello_world@",
			suffix:    "ld@",
			wantMatch: true,
		},
		{
			name:      "unicode_characters",
			input:     "こんにちは世界",
			suffix:    "世界",
			wantMatch: true,
		},
		{
			name:      "unicode_no_match",
			input:     "こんにちは世界",
			suffix:    "こんにちは",
			wantMatch: false,
		},
	}

	goldie := newGoldie(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := match.StringLikeHasSuffixTm(tm, tt.suffix)
			matched, explanation := matcher.Match(tt.input)

			if matched != tt.wantMatch {
				t.Errorf("StringLikeHasSuffixTm().Match() matched = %v, want %v", matched, tt.wantMatch)
			}

			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}

func TestStringLikeHasSuffix(t *testing.T) {
	type MyString string

	tests := []struct {
		name      string
		input     MyString
		suffix    string
		wantMatch bool
	}{
		{
			name:      "ends_with_suffix",
			input:     "hello world",
			suffix:    "world",
			wantMatch: true,
		},
		{
			name:      "does_not_end_with_suffix",
			input:     "hello world",
			suffix:    "hello",
			wantMatch: false,
		},
		{
			name:      "empty_suffix",
			input:     "test",
			suffix:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_suffix",
			input:     "",
			suffix:    "test",
			wantMatch: false,
		},
		{
			name:      "both_empty",
			input:     "",
			suffix:    "",
			wantMatch: true,
		},
		{
			name:      "special_characters",
			input:     "hello@world.com",
			suffix:    ".com",
			wantMatch: true,
		},
		{
			name:      "unicode_characters",
			input:     "こんにちは世界",
			suffix:    "世界",
			wantMatch: true,
		},
		{
			name:      "partial_unicode_no_match",
			input:     "こんにちは世界",
			suffix:    "こんにちは",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_uppercase",
			input:     "Hello World",
			suffix:    "WORLD",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_correct_case",
			input:     "Hello World",
			suffix:    "World",
			wantMatch: true,
		},
		{
			name:      "numeric_suffix",
			input:     "abc123",
			suffix:    "123",
			wantMatch: true,
		},
		{
			name:      "mixed_alphanumeric",
			input:     "abc123def",
			suffix:    "3def",
			wantMatch: true,
		},
	}

	goldie := newGoldie(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := match.StringLikeHasSuffix[MyString](tt.suffix)
			matched, explanation := matcher.Match(tt.input)

			if matched != tt.wantMatch {
				t.Errorf("StringLikeHasSuffix().Match() matched = %v, want %v", matched, tt.wantMatch)
			}

			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}

func TestStringHasSuffix(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		suffix    string
		wantMatch bool
	}{
		{
			name:      "simple_ends_with",
			input:     "hello world",
			suffix:    "world",
			wantMatch: true,
		},
		{
			name:      "does_not_end_with",
			input:     "hello world",
			suffix:    "hello",
			wantMatch: false,
		},
		{
			name:      "empty_suffix_always_matches",
			input:     "anything",
			suffix:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_empty_suffix",
			input:     "",
			suffix:    "",
			wantMatch: true,
		},
		{
			name:      "empty_string_with_non_empty_suffix",
			input:     "",
			suffix:    "test",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_different_case",
			input:     "Hello World",
			suffix:    "WORLD",
			wantMatch: false,
		},
		{
			name:      "case_sensitive_same_case",
			input:     "Hello World",
			suffix:    "World",
			wantMatch: true,
		},
		{
			name:      "suffix_longer_than_input",
			input:     "hi",
			suffix:    "hello",
			wantMatch: false,
		},
		{
			name:      "exact_match",
			input:     "exact",
			suffix:    "exact",
			wantMatch: true,
		},
		{
			name:      "whitespace_handling",
			input:     "hello world ",
			suffix:    " ",
			wantMatch: true,
		},
		{
			name:      "newline_characters",
			input:     "line1\n",
			suffix:    "\n",
			wantMatch: true,
		},
		{
			name:      "tab_characters",
			input:     "col1\t",
			suffix:    "\t",
			wantMatch: true,
		},
		{
			name:      "single_character_match",
			input:     "test",
			suffix:    "t",
			wantMatch: true,
		},
		{
			name:      "single_character_no_match",
			input:     "test",
			suffix:    "x",
			wantMatch: false,
		},
		{
			name:      "numeric_string",
			input:     "12345",
			suffix:    "345",
			wantMatch: true,
		},
		{
			name:      "mixed_content",
			input:     "abc123def",
			suffix:    "def",
			wantMatch: true,
		},
		{
			name:      "repeated_characters",
			input:     "aaabbb",
			suffix:    "bbb",
			wantMatch: true,
		},
		{
			name:      "partial_repeated_characters",
			input:     "aaabbb",
			suffix:    "bb",
			wantMatch: true,
		},
		{
			name:      "url_like_string",
			input:     "https://example.com",
			suffix:    ".com",
			wantMatch: true,
		},
		{
			name:      "path_like_string",
			input:     "/usr/local/bin",
			suffix:    "/bin",
			wantMatch: true,
		},
		{
			name:      "file_extension",
			input:     "document.pdf",
			suffix:    ".pdf",
			wantMatch: true,
		},
		{
			name:      "wrong_file_extension",
			input:     "document.pdf",
			suffix:    ".txt",
			wantMatch: false,
		},
	}

	goldie := newGoldie(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := match.StringHasSuffix(tt.suffix)
			matched, explanation := matcher.Match(tt.input)

			if matched != tt.wantMatch {
				t.Errorf("StringHasSuffix().Match() matched = %v, want %v", matched, tt.wantMatch)
			}

			goldie.Assert(t, tt.name, []byte(explanation))
		})
	}
}
