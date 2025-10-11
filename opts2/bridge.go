package opts2

type BridgeOne func(Opts, Vals, Func) Out

func (b BridgeOne) Match(f Func) Func {
	return func(o Opts, v Vals) Out {
		return b(o, v, f)
	}
}

type BridgeMany func(Opts, Vals, ...Func) Out

func (b BridgeMany) Match(fs ...Func) Func {
	return func(o Opts, v Vals) Out {
		return b(o, v, fs...)
	}
}
