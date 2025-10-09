package push3

import (
	"fmt"
	"iter"

	"github.com/krelinga/go-match/matchfmt"
)

type labeledResult struct {
	label  string
	result Result
}

type ResultBuilder struct {
	results []labeledResult
}

func (rb *ResultBuilder) Add(label string, result Result) {
	rb.results = append(rb.results, labeledResult{label: label, result: result})
}

func (rb *ResultBuilder) All() iter.Seq[bool] {
	return func(yield func(bool) bool) {
		for _, lr := range rb.results {
			if !yield(lr.result.Matched) {
				return
			}
		}
	}
}

func (rb *ResultBuilder) Finish(match bool, headline string) Result {
	details := make([]string, len(rb.results))
	for i, lr := range rb.results {
		details[i] = fmt.Sprintf("%s:\n%s", lr.label, matchfmt.Indent(lr.result.Explanation))
	}
	return Result{
		Matched:     match,
		Explanation: matchfmt.Explain(match, headline, details...),
	}
}

func All(in iter.Seq[bool]) bool {
	for b := range in {
		if !b {
			return false
		}
	}
	return true
}

func Any(in iter.Seq[bool]) bool {
	for b := range in {
		if b {
			return true
		}
	}
	return false
}
