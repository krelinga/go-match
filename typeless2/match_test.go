package typeless2_test

import (
	"errors"
	"testing"

	"github.com/krelinga/go-match/typeless2"
)

type Person struct {
	Name string
	Age  int
}

func match(t *testing.T, in typeless2.Input, m typeless2.Func) bool {
	t.Helper()
	out := m(in)
	if out.Err != nil {
		t.Errorf("unexpected error: %v", out.Err)
		return false
	}
	if !out.Matched {
		t.Errorf("expected match, but got:\n%s", out.Explanation)
	}
	return out.Matched
}

func NoError() (int, error) {
	return 42, nil
}

func WithError() (int, error) {
	return 0, errors.New("some error")
}

func TestFoo(t *testing.T) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	match(t, typeless2.NewInput(people[0]),
		typeless2.AllOf(
			typeless2.Field("Name", typeless2.Equal("Alice")),
			typeless2.Field("Age", typeless2.Equal(30)),
		),
	)
	match(t, typeless2.NewInput(WithError()),
		typeless2.NoError(
			typeless2.Equal(42),
		),
	)
	match(t, typeless2.NewInput(NoError()),
		typeless2.NoError(
			typeless2.Equal(42),
		),
	)
	match(t, typeless2.NewInput(42, 10),
		typeless2.NoError(
			typeless2.Equal(42),
		),
	)
}