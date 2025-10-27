package opts3_test

import (
	"testing"

	"github.com/krelinga/go-match/opts3"
)

func TestNewResult(t *testing.T) {
	tests := []struct {
		name     string
		code     opts3.Code
		str      string
		wantCode opts3.Code
		wantStr  string
	}{
		{
			name:     "create result with Err code",
			code:     opts3.Err,
			str:      "error message",
			wantCode: opts3.Err,
			wantStr:  "error message",
		},
		{
			name:     "create result with Yes code",
			code:     opts3.Yes,
			str:      "success message",
			wantCode: opts3.Yes,
			wantStr:  "success message",
		},
		{
			name:     "create result with No code",
			code:     opts3.No,
			str:      "failure message",
			wantCode: opts3.No,
			wantStr:  "failure message",
		},
		{
			name:     "create result with empty string",
			code:     opts3.Yes,
			str:      "",
			wantCode: opts3.Yes,
			wantStr:  "",
		},
		{
			name:     "create result with multiline string",
			code:     opts3.Err,
			str:      "line 1\nline 2\nline 3",
			wantCode: opts3.Err,
			wantStr:  "line 1\nline 2\nline 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := opts3.NewResult(tt.code, tt.str)

			gotCode := result.Code()
			if gotCode != tt.wantCode {
				t.Errorf("NewResult().Code() = %v, want %v", gotCode, tt.wantCode)
			}

			gotStr := result.String()
			if gotStr != tt.wantStr {
				t.Errorf("NewResult().String() = %q, want %q", gotStr, tt.wantStr)
			}
		})
	}
}

func TestNewResultPanic(t *testing.T) {
	tests := []struct {
		name string
		code opts3.Code
		str  string
	}{
		{
			name: "invalid code panics",
			code: opts3.Code(99),
			str:  "test message",
		},
		{
			name: "negative code panics",
			code: opts3.Code(-1),
			str:  "test message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("NewResult() did not panic for invalid code %v", tt.code)
				}
			}()
			_ = opts3.NewResult(tt.code, tt.str)
		})
	}
}
