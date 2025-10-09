package typeless2

import "github.com/krelinga/go-match/matchfmt"

func AllOf(funcs ...Func) Func {
	return func(in Input) Output {
		childOutputs := make([]Output, len(funcs))
		for i, f := range funcs {
			childOutputs[i] = f(in)
			if childOutputs[i].Err != nil {
				return NewOutputErr(childOutputs[i].Err)
			}
		}
		matched := true
		childExplanations := make([]string, len(childOutputs))
		for i, out := range childOutputs {
			matched = matched && out.Matched
			childExplanations[i] = out.Explanation
		}
		return Output{
			Matched: matched,
			Explanation: matchfmt.Explain(matched, "AllOf", childExplanations...),
		}
	}
}