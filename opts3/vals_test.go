package opts3_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/krelinga/go-match/opts3"
)

func TestNewVals(t *testing.T) {
	ret1 := func() int {
		return 42
	}
	ret2 := func() (string, float64) {
		return "hello", 3.14
	}
	tests := []struct {
		name string
		f    func() opts3.Vals
		wantTypes []reflect.Type
	}{
		{
			name: "mixed types",
			f: func() opts3.Vals {
				return opts3.NewVals(42, "hello", 3.14)
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[int](),
				reflect.TypeFor[string](),
				reflect.TypeFor[float64](),
			},
		},
		{
			name: "single return",
			f: func() opts3.Vals {
				return opts3.NewVals1(ret1())
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[int](),
			},
		},
		{
			name: "two returns",
			f: func() opts3.Vals {
				return opts3.NewVals2(ret2())
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[string](),
				reflect.TypeFor[float64](),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vals := tt.f()
			if len(vals) != len(tt.wantTypes) {
				t.Fatalf("got %d vals, want %d", len(vals), len(tt.wantTypes))
			}
			for i, wantType := range tt.wantTypes {
				if vals[i].Type() != wantType {
					t.Errorf("val %d: got type %v, want %v", i, vals[i].Type(), wantType)
				}
			}
		})
	}
}

func TestNewVals1(t *testing.T) {
	retInt := func() int {
		return 42
	}
	retErr := func() error {
		return errors.New("test error")
	}
	retNilErr := func() error {
		return nil
	}
	tests := []struct {
		name    string
		f      func() opts3.Vals
		wantType reflect.Type
	}{
		{
			name: "int",
			f: func() opts3.Vals {
				return opts3.NewVals1(retInt())
			},
			wantType: reflect.TypeFor[int](),
		},
		{
			name: "error",
			f: func() opts3.Vals {
				return opts3.NewVals1(retErr())
			},
			wantType: reflect.TypeFor[error](),
		},
		{
			name: "nil error",
			f: func() opts3.Vals {
				return opts3.NewVals1(retNilErr())
			},
			wantType: reflect.TypeFor[error](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vals := tt.f()
			if len(vals) != 1 {
				t.Fatalf("got %d vals, want 1", len(vals))
			}
			if vals[0].Type() != tt.wantType {
				t.Errorf("got type %v, want %v", vals[0].Type(), tt.wantType)
			}
		})
	}
}

func TestNewVals2(t *testing.T) {
	retStrFloat := func() (string, float64) {
		return "hello", 3.14
	}
	retBoolInt := func() (bool, int) {
		return true, 7
	}
	retIntErr := func() (int, error) {
		return 0, nil
	}
	retIntNilErr := func() (int, error) {
		return 0, nil
	}
	tests := []struct {
		name     string
		f       func() opts3.Vals
		wantTypes []reflect.Type
	}{
		{
			name: "string and float64",
			f: func() opts3.Vals {
				return opts3.NewVals2(retStrFloat())
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[string](),
				reflect.TypeFor[float64](),
			},
		},
		{
			name: "bool and int",
			f: func() opts3.Vals {
				return opts3.NewVals2(retBoolInt())
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[bool](),
				reflect.TypeFor[int](),
			},
		},
		{
			name: "int and error",
			f: func() opts3.Vals {
				return opts3.NewVals2(retIntErr())
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[int](),
				reflect.TypeFor[error](),
			},
		},
		{
			name: "int and nil error",
			f: func() opts3.Vals {
				return opts3.NewVals2(retIntNilErr())
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[int](),
				reflect.TypeFor[error](),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vals := tt.f()
			if len(vals) != len(tt.wantTypes) {
				t.Fatalf("got %d vals, want %d", len(vals), len(tt.wantTypes))
			}
			for i, wantType := range tt.wantTypes {
				if vals[i].Type() != wantType {
					t.Errorf("val %d: got type %v, want %v", i, vals[i].Type(), wantType)
				}
			}
		})
	}
}

func TestNewVals3(t *testing.T) {
	retStrFloatBool := func() (string, float64, bool) {
		return "hello", 3.14, true
	}
	retIntIntErr := func() (int, int, error) {
		return 1, 2, errors.New("test error")
	}
	retIntIntNilErr := func() (int, int, error) {
		return 1, 2, nil
	}
	tests := []struct {
		name     string
		f       func() opts3.Vals
		wantTypes []reflect.Type
	}{
		{
			name: "string, float64, bool",
			f: func() opts3.Vals {
				return opts3.NewVals3(retStrFloatBool())
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[string](),
				reflect.TypeFor[float64](),
				reflect.TypeFor[bool](),
			},
		},
		{
			name: "int, int, error",
			f: func() opts3.Vals {
				return opts3.NewVals3(retIntIntErr())
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[int](),
				reflect.TypeFor[int](),
				reflect.TypeFor[error](),
			},
		},
		{
			name: "int, int, nil error",
			f: func() opts3.Vals {
				return opts3.NewVals3(retIntIntNilErr())
			},
			wantTypes: []reflect.Type{
				reflect.TypeFor[int](),
				reflect.TypeFor[int](),
				reflect.TypeFor[error](),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vals := tt.f()
			if len(vals) != len(tt.wantTypes) {
				t.Fatalf("got %d vals, want %d", len(vals), len(tt.wantTypes))
			}
			for i, wantType := range tt.wantTypes {
				if vals[i].Type() != wantType {
					t.Errorf("val %d: got type %v, want %v", i, vals[i].Type(), wantType)
				}
			}
		})
	}
}

func TestWant1(t *testing.T) {
	tests := []struct {
		name    string
		vals    opts3.Vals
		f       func(*testing.T, opts3.Vals)
	}{
		{
			name: "correct type",
			vals: opts3.NewVals1(42),
			f: func(t *testing.T, vals opts3.Vals) {
				got, err := opts3.Want1[int](vals)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				want := 42
				if got != want {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			name: "incorrect type",
			vals: opts3.NewVals1("hello"),
			f: func(t *testing.T, vals opts3.Vals) {
				_, err := opts3.Want1[int](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		}, {
			name: "invalid value",
			vals: opts3.Vals{reflect.Value{}},
			f: func(t *testing.T, vals opts3.Vals) {
				_, err := opts3.Want1[int](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		}, {
			name: "wrong number of values",
			vals: opts3.NewVals2(42, "extra"),
			f: func(t *testing.T, vals opts3.Vals) {
				_, err := opts3.Want1[int](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t, tt.vals)
		})
	}
}

func TestWant2(t *testing.T) {
	tests := []struct {
		name    string
		vals    opts3.Vals
		f       func(*testing.T, opts3.Vals)
	}{
		{
			name: "correct types",
			vals: opts3.NewVals2(42, "hello"),
			f: func(t *testing.T, vals opts3.Vals) {
				got1, got2, err := opts3.Want2[int, string](vals)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				want1 := 42
				want2 := "hello"
				if got1 != want1 {
					t.Errorf("got1 %v, want %v", got1, want1)
				}
				if got2 != want2 {
					t.Errorf("got2 %v, want %v", got2, want2)
				}
			},
		},
		{
			name: "incorrect first type",
			vals: opts3.NewVals2("hello", 42),
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, err := opts3.Want2[int, int](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
		{
			name: "incorrect second type",
			vals: opts3.NewVals2(42, 3.14),
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, err := opts3.Want2[int, string](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
		{
			name: "invalid value",
			vals: opts3.Vals{reflect.Value{}, reflect.ValueOf("hello")},
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, err := opts3.Want2[int, string](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
		{
			name: "wrong number of values - too few",
			vals: opts3.NewVals1(42),
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, err := opts3.Want2[int, string](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
		{
			name: "wrong number of values - too many",
			vals: opts3.NewVals3(42, "hello", 3.14),
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, err := opts3.Want2[int, string](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t, tt.vals)
		})
	}
}

func TestWant3(t *testing.T) {
	tests := []struct {
		name    string
		vals    opts3.Vals
		f       func(*testing.T, opts3.Vals)
	}{
		{
			name: "correct types",
			vals: opts3.NewVals3(42, "hello", 3.14),
			f: func(t *testing.T, vals opts3.Vals) {
				got1, got2, got3, err := opts3.Want3[int, string, float64](vals)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				want1 := 42
				want2 := "hello"
				want3 := 3.14
				if got1 != want1 {
					t.Errorf("got1 %v, want %v", got1, want1)
				}
				if got2 != want2 {
					t.Errorf("got2 %v, want %v", got2, want2)
				}
				if got3 != want3 {
					t.Errorf("got3 %v, want %v", got3, want3)
				}
			},
		},
		{
			name: "incorrect first type",
			vals: opts3.NewVals3("hello", 42, 3.14),
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, _, err := opts3.Want3[int, int, float64](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
		{
			name: "incorrect second type",
			vals: opts3.NewVals3(42, 3.14, "hello"),
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, _, err := opts3.Want3[int, string, string](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
		{
			name: "incorrect third type",
			vals: opts3.NewVals3(42, "hello", true),
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, _, err := opts3.Want3[int, string, float64](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
		{
			name: "invalid value",
			vals: opts3.Vals{reflect.Value{}, reflect.ValueOf("hello"), reflect.ValueOf(3.14)},
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, _, err := opts3.Want3[int, string, float64](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
		{
			name: "wrong number of values - too few",
			vals: opts3.NewVals2(42, "hello"),
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, _, err := opts3.Want3[int, string, float64](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
		{
			name: "wrong number of values - too many",
			vals: opts3.NewVals(42, "hello", 3.14, true),
			f: func(t *testing.T, vals opts3.Vals) {
				_, _, _, err := opts3.Want3[int, string, float64](vals)
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t, tt.vals)
		})
	}
}