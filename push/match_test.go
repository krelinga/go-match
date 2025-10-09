package push_test

import (
	"testing"
	"time"

	"github.com/krelinga/go-match/push"
)

func match[T any](t *testing.T, got T, matcher push.M[T]) {
	t.Helper()
	res := matcher(got)
	if !res.Matched {
		t.Errorf("expected match for %v, but got %v", got, res.Reason)
	}
}

type Person struct {
	Name   string
	Age    int
	Joined time.Time
	Tags   []string
}

func TestBasicEquality(t *testing.T) {
	match(t, 42, func(got int) push.Result {
		return push.Results{
			push.Match(got, push.Equal(42)),
			push.Match(got, push.EqualOpts[int]{}.M),
		}.AllOf()
	})
}
