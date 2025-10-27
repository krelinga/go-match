package opts3_test

import (
	"testing"

	"github.com/krelinga/go-match/opts3"
)

func TestCodePanicIfInvalid(t *testing.T) {
	tests := []struct {
		name      string
		code      opts3.Code
		shouldPanic bool
	}{
		{
			name:        "Err is valid",
			code:        opts3.Err,
			shouldPanic: false,
		},
		{
			name:        "Yes is valid",
			code:        opts3.Yes,
			shouldPanic: false,
		},
		{
			name:        "No is valid",
			code:        opts3.No,
			shouldPanic: false,
		},
		{
			name:        "unknown code value panics",
			code:        opts3.Code(99),
			shouldPanic: true,
		},
		{
			name:        "negative code value panics",
			code:        opts3.Code(-1),
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if tt.shouldPanic && r == nil {
					t.Errorf("PanicIfInvalid() should have panicked for %v", tt.code)
				}
				if !tt.shouldPanic && r != nil {
					t.Errorf("PanicIfInvalid() should not have panicked for %v, got panic: %v", tt.code, r)
				}
			}()
			tt.code.PanicIfInvalid()
		})
	}
}

func TestCodeString(t *testing.T) {
	tests := []struct {
		name string
		code opts3.Code
		want string
	}{
		{
			name: "Err returns Err string",
			code: opts3.Err,
			want: "Err",
		},
		{
			name: "Yes returns Yes string",
			code: opts3.Yes,
			want: "Yes",
		},
		{
			name: "No returns No string",
			code: opts3.No,
			want: "No",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.code.String()
			if got != tt.want {
				t.Errorf("Code.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCodeStringPanic(t *testing.T) {
	tests := []struct {
		name string
		code opts3.Code
	}{
		{
			name: "unknown code value panics",
			code: opts3.Code(99),
		},
		{
			name: "negative code value panics",
			code: opts3.Code(-1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Code.String() did not panic for unknown value %d", tt.code)
				}
			}()
			_ = tt.code.String()
		})
	}
}

func TestCodeEmoji(t *testing.T) {
	tests := []struct {
		name string
		code opts3.Code
		want string
	}{
		{
			name: "Err returns warning emoji",
			code: opts3.Err,
			want: "⚠️",
		},
		{
			name: "Yes returns check mark emoji",
			code: opts3.Yes,
			want: "✅",
		},
		{
			name: "No returns cross mark emoji",
			code: opts3.No,
			want: "❌",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.code.Emoji()
			if got != tt.want {
				t.Errorf("Code.Emoji() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCodeEmojiPanic(t *testing.T) {
	tests := []struct {
		name string
		code opts3.Code
	}{
		{
			name: "unknown code value panics",
			code: opts3.Code(99),
		},
		{
			name: "negative code value panics",
			code: opts3.Code(-1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Code.Emoji() did not panic for unknown value %d", tt.code)
				}
			}()
			_ = tt.code.Emoji()
		})
	}
}

func TestCodeNot(t *testing.T) {
	tests := []struct {
		name string
		code opts3.Code
		want opts3.Code
	}{
		{
			name: "Not of Err returns Err",
			code: opts3.Err,
			want: opts3.Err,
		},
		{
			name: "Not of Yes returns No",
			code: opts3.Yes,
			want: opts3.No,
		},
		{
			name: "Not of No returns Yes",
			code: opts3.No,
			want: opts3.Yes,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := opts3.CodeNot(tt.code)
			if got != tt.want {
				t.Errorf("CodeNot(%v) = %v, want %v", tt.code, got, tt.want)
			}
		})
	}
}

func TestCodeNotPanic(t *testing.T) {
	tests := []struct {
		name string
		code opts3.Code
	}{
		{
			name: "unknown code value panics",
			code: opts3.Code(99),
		},
		{
			name: "negative code value panics",
			code: opts3.Code(-1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("CodeNot() did not panic for unknown value %d", tt.code)
				}
			}()
			_ = opts3.CodeNot(tt.code)
		})
	}
}

func TestCodeAnd(t *testing.T) {
	tests := []struct {
		name  string
		codes []opts3.Code
		want  opts3.Code
	}{
		{
			name:  "empty slice returns Err",
			codes: []opts3.Code{},
			want:  opts3.Err,
		},
		{
			name:  "single Yes returns Yes",
			codes: []opts3.Code{opts3.Yes},
			want:  opts3.Yes,
		},
		{
			name:  "single No returns No",
			codes: []opts3.Code{opts3.No},
			want:  opts3.No,
		},
		{
			name:  "single Err returns Err",
			codes: []opts3.Code{opts3.Err},
			want:  opts3.Err,
		},
		{
			name:  "all Yes returns Yes",
			codes: []opts3.Code{opts3.Yes, opts3.Yes, opts3.Yes},
			want:  opts3.Yes,
		},
		{
			name:  "any No returns No when no Err",
			codes: []opts3.Code{opts3.Yes, opts3.No, opts3.Yes},
			want:  opts3.No,
		},
		{
			name:  "any Err returns Err regardless of position",
			codes: []opts3.Code{opts3.Yes, opts3.Err, opts3.Yes},
			want:  opts3.Err,
		},
		{
			name:  "Err takes precedence over No",
			codes: []opts3.Code{opts3.No, opts3.Err},
			want:  opts3.Err,
		},
		{
			name:  "first Err encountered returns Err",
			codes: []opts3.Code{opts3.Yes, opts3.Err, opts3.No},
			want:  opts3.Err,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := opts3.CodeAnd(tt.codes...)
			if got != tt.want {
				t.Errorf("CodeAnd(%v) = %v, want %v", tt.codes, got, tt.want)
			}
		})
	}
}

func TestCodeAndPanic(t *testing.T) {
	tests := []struct {
		name  string
		codes []opts3.Code
	}{
		{
			name:  "invalid code in slice panics",
			codes: []opts3.Code{opts3.Yes, opts3.Code(99), opts3.No},
		},
		{
			name:  "first invalid code panics",
			codes: []opts3.Code{opts3.Code(-1)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("CodeAnd() did not panic for invalid codes %v", tt.codes)
				}
			}()
			_ = opts3.CodeAnd(tt.codes...)
		})
	}
}

func TestCodeOr(t *testing.T) {
	tests := []struct {
		name  string
		codes []opts3.Code
		want  opts3.Code
	}{
		{
			name:  "empty slice returns Err",
			codes: []opts3.Code{},
			want:  opts3.Err,
		},
		{
			name:  "single Yes returns Yes",
			codes: []opts3.Code{opts3.Yes},
			want:  opts3.Yes,
		},
		{
			name:  "single No returns No",
			codes: []opts3.Code{opts3.No},
			want:  opts3.No,
		},
		{
			name:  "single Err returns Err",
			codes: []opts3.Code{opts3.Err},
			want:  opts3.Err,
		},
		{
			name:  "all No returns No",
			codes: []opts3.Code{opts3.No, opts3.No, opts3.No},
			want:  opts3.No,
		},
		{
			name:  "any Yes returns Yes when no Err",
			codes: []opts3.Code{opts3.No, opts3.Yes, opts3.No},
			want:  opts3.Yes,
		},
		{
			name:  "any Err returns Err regardless of position",
			codes: []opts3.Code{opts3.No, opts3.Err, opts3.No},
			want:  opts3.Err,
		},
		{
			name:  "Err takes precedence over Yes",
			codes: []opts3.Code{opts3.Yes, opts3.Err},
			want:  opts3.Err,
		},
		{
			name:  "first Err encountered returns Err",
			codes: []opts3.Code{opts3.No, opts3.Err, opts3.Yes},
			want:  opts3.Err,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := opts3.CodeOr(tt.codes...)
			if got != tt.want {
				t.Errorf("CodeOr(%v) = %v, want %v", tt.codes, got, tt.want)
			}
		})
	}
}

func TestCodeOrPanic(t *testing.T) {
	tests := []struct {
		name  string
		codes []opts3.Code
	}{
		{
			name:  "invalid code in slice panics",
			codes: []opts3.Code{opts3.No, opts3.Code(99), opts3.Yes},
		},
		{
			name:  "first invalid code panics",
			codes: []opts3.Code{opts3.Code(-1)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("CodeOr() did not panic for invalid codes %v", tt.codes)
				}
			}()
			_ = opts3.CodeOr(tt.codes...)
		})
	}
}
