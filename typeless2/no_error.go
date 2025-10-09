package typeless2

import (
	"fmt"
)

func NoError(f Func) Func {
	return func(in Input) Output {
		if len(in) == 0 {
			return Output{
				Err: fmt.Errorf("%w: no input values", ErrInputLen),
			}
		}
		last := in[len(in)-1]
		if last == nil {
			return f(in[:len(in)-1])
		}
		asErr, ok := last.(error)
		if !ok {
			return Output{
				Err: fmt.Errorf("%w: last input value is not an error, got %T", ErrType, last),
			}
		}
		return Output{
			Matched:     false,
			Explanation: fmt.Sprintf("input has error: %v", asErr),
		}
	}
}
