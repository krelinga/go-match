package typeless2

type Output struct {
	Matched bool
	Explanation string
	Err error
}

func NewOutputErr(err error) Output {
	return Output{
		Err: err,
	}
}