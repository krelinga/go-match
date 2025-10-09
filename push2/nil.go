package push2

import "github.com/krelinga/go-match/matchfmt"

func MatchNil(isNil bool, matcher Matcher[bool]) Result {
	r := matcher.Match(isNil)
	return Result{
		Matched: r.Matched,
		Reason: matchfmt.Explain(r.Matched, "Nil", r.Reason),
	}	
}