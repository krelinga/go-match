package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchutil"
)

func AllOf[T any](matchers ...Matcher[T]) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = true
		details := make([]string, 0, len(matchers))
		for i, matcher := range matchers {
			m, e := matcher.Match(got)
			if !m {
				matched = false
			}
			indexDetail := fmt.Sprintf("matcher %d:", i)
			matcherDetail := matchutil.Indent(e)
			details = append(details, indexDetail, matcherDetail)
		}
		explanation = matchutil.Explain(matched, "match.AllOf", details...)
		return
	})
}

func AnyOf[T any](matchers ...Matcher[T]) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = false
		details := make([]string, 0, len(matchers))
		for i, matcher := range matchers {
			m, e := matcher.Match(got)
			if m {
				matched = true
			}
			indexDetail := fmt.Sprintf("matcher %d:", i)
			matcherDetail := matchutil.Indent(e)
			details = append(details, indexDetail, matcherDetail)
		}
		explanation = matchutil.Explain(matched, "match.AnyOf", details...)
		return
	})
}

func Not[T any](matcher Matcher[T]) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		m, e := matcher.Match(got)
		matched = !m
		explanation = matchutil.Explain(matched, "match.Not", "negated matcher:", matchutil.Indent(e))
		return
	})
}

func Alway[T any]() Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = true
		explanation = matchutil.Explain(matched, "match.Alway", "always matches")
		return
	})
}

func Never[T any]() Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = false
		explanation = matchutil.Explain(matched, "match.Never", "never matches")
		return
	})
}
