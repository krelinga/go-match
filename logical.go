package match

import (
	"fmt"

	"github.com/krelinga/go-match/matchfmt"
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
			matcherDetail := matchfmt.Indent(e)
			details = append(details, indexDetail, matcherDetail)
		}
		explanation = matchfmt.Explain(matched, "match.AllOf", details...)
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
			matcherDetail := matchfmt.Indent(e)
			details = append(details, indexDetail, matcherDetail)
		}
		explanation = matchfmt.Explain(matched, "match.AnyOf", details...)
		return
	})
}

func Not[T any](matcher Matcher[T]) Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		m, e := matcher.Match(got)
		matched = !m
		explanation = matchfmt.Explain(matched, "match.Not", "negated matcher:", matchfmt.Indent(e))
		return
	})
}

func Alway[T any]() Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = true
		explanation = matchfmt.Explain(matched, "match.Alway", "always matches")
		return
	})
}

func Never[T any]() Matcher[T] {
	return MatcherFunc[T](func(got T) (matched bool, explanation string) {
		matched = false
		explanation = matchfmt.Explain(matched, "match.Never", "never matches")
		return
	})
}
