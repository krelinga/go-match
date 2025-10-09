package push2

import "github.com/krelinga/go-match/matchfmt"

type Result struct {
	Matched bool
	Reason  string
}

func AllOfResults(results ...Result) Result {
	reasons := make([]string, len(results))
	matched := true
	for i, r := range results {
		reasons[i] = r.Reason
		matched = matched && r.Matched
	}
	return Result{
		Matched: matched,
		Reason:  matchfmt.Explain(matched, "AllOf", reasons...),
	}
}