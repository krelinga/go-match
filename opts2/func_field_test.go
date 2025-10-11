package opts2_test

import (
	"testing"

	"github.com/krelinga/go-match/opts2"
)

func match(t *testing.T, vals opts2.Vals, f opts2.Func) {
	opts := opts2.NewOpts()
	out := f(opts, vals)
	if out.Err != nil {
		t.Errorf("unexpected error: %v", out.Err)
	} else if !out.Matched {
		t.Errorf("not matched: %s", out.Note)
	}
}

type Person struct {
	Name string
	Age  int
}

func TestField(t *testing.T) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	match(t, opts2.NewVals(people[0]),
		opts2.Field("Name").Match(opts2.Equal("Alice")),
	)
}
