package match

import (
	"cmp"
	"fmt"

	"github.com/krelinga/go-match/matchutil"
)

func Equal[T comparable](want T) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = got == want
		expected := fmt.Sprintf("got == %s", matchutil.DefaultFormat(want))
		var info string
		if !matched {
			actual := fmt.Sprintf("got == %s", matchutil.DefaultFormat(got))
			info = matchutil.ActualVsExpected(actual, expected)
		} else {
			info = expected
		}
		explanation = matchutil.Explain(matched, "match.Equal", info)
		return
	})
}

func NotEqual[T comparable](other T) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = got != other
		expected := fmt.Sprintf("got != %s", matchutil.DefaultFormat(other))
		var info string
		if !matched {
			actual := fmt.Sprintf("got == %s", matchutil.DefaultFormat(got))
			info = matchutil.ActualVsExpected(actual, expected)
		} else {
			info = expected
		}
		explanation = matchutil.Explain(matched, "match.NotEqual", info)
		return
	})
}

func LessThan[T cmp.Ordered](other T) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = got < other
		expected := fmt.Sprintf("got < %s", matchutil.DefaultFormat(other))
		var info string
		if !matched {
			actual := fmt.Sprintf("got == %s", matchutil.DefaultFormat(got))
			info = matchutil.ActualVsExpected(actual, expected)
		} else {
			info = expected
		}
		explanation = matchutil.Explain(matched, "match.LessThan", info)
		return
	})
}

func LessThanOrEqual[T cmp.Ordered](other T) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = got <= other
		expected := fmt.Sprintf("got <= %s", matchutil.DefaultFormat(other))
		var info string
		if !matched {
			actual := fmt.Sprintf("got == %s", matchutil.DefaultFormat(got))
			info = matchutil.ActualVsExpected(actual, expected)
		} else {
			info = expected
		}
		explanation = matchutil.Explain(matched, "match.LessThanOrEqual", info)
		return
	})
}

func GreaterThan[T cmp.Ordered](other T) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = got > other
		expected := fmt.Sprintf("got > %s", matchutil.DefaultFormat(other))
		var info string
		if !matched {
			actual := fmt.Sprintf("got == %s", matchutil.DefaultFormat(got))
			info = matchutil.ActualVsExpected(actual, expected)
		} else {
			info = expected
		}
		explanation = matchutil.Explain(matched, "match.GreaterThan", info)
		return
	})
}

func GreaterThanOrEqual[T cmp.Ordered](other T) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = got >= other
		expected := fmt.Sprintf("got >= %s", matchutil.DefaultFormat(other))
		var info string
		if !matched {
			actual := fmt.Sprintf("got == %s", matchutil.DefaultFormat(got))
			info = matchutil.ActualVsExpected(actual, expected)
		} else {
			info = expected
		}
		explanation = matchutil.Explain(matched, "match.GreaterThanOrEqual", info)
		return
	})
}
