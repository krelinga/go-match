package opts2

func All() BridgeMany {
	return func(o Opts, v Vals, fs ...Func) Out {
		matched := true
		for _, f := range fs {
			out := f(o, v)
			if out.Err != nil {
				return Out{Err: out.Err}
			}
			if !out.Matched {
				matched = false
			}
		}
		return Out{Matched: matched}
	}
}