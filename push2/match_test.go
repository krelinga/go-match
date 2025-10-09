package push2_test

import (
	"testing"

	"github.com/krelinga/go-match/push2"
)

type Person struct {
	Name string
	Age  int
}

func match[T any](t *testing.T, got T, fn push2.Func[T]) {
	t.Helper()
	res := fn.Match(got)
	if !res.Matched {
		t.Errorf("expected match for %v, but got\n%v", got, res.Reason)
	}
}

func TestFoo(t *testing.T) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	match(t, people, func(got []Person) push2.Result {
		return push2.AllOfResults(
			push2.MatchNil(got == nil, push2.Equal(false)),
			push2.MatchLen(len(got), push2.Equal(2)),
			push2.Match(got[0].Name, push2.Equal("Alice")),
			push2.Match(got[1].Age, push2.EqualMatcher[int]{
				Val: 25,
				Fmt: push2.DefaultFmtFunc[int](),
				Cmp: push2.DefaultCmp[int](),
			}),
		)
	})
	match(t, people, func(got []Person) push2.Result {
		return push2.MatchNil(got == nil, push2.Equal(false))
	})
}
