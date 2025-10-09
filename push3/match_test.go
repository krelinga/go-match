package push3_test

import (
	"testing"

	"github.com/krelinga/go-match/push3"
)

type Person struct {
	Name string
	Age  int
}

func matched(t *testing.T, r push3.Result) {
	t.Helper()
	if !r.Matched {
		t.Errorf("expected match, but got:\n%s", r.Explanation)
	}
}

func TestFoo(t *testing.T) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	matched(t, push3.MatchLen(push3.WrapSlice(people), push3.EqualM(2)))
	matched(t, push3.MatchNil(push3.WrapSlice(people), push3.EqualM(false)))
	matched(t, push3.EqualM(10)(10))
	matched(t, push3.Match(10, push3.EqualM(10)))
	matchPerson := func(wantName string, wantAge int) push3.M[Person] {
		return func(got Person) push3.Result {
			b := push3.ResultBuilder{}
			b.Add("Name", push3.Match(got.Name, push3.EqualM(wantName)))
			b.Add("Age", push3.Match(got.Age, push3.EqualM(wantAge)))
			return b.Finish(push3.All(b.All()), "Person")
		}
	}
	matched(t, push3.Match(people[0], matchPerson("Alice", 31)))
}
