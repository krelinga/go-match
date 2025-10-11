package opts2

type Bridge func(Opts, Vals, Func) Out

func (b Bridge) Match(f Func) Func {
	return func(o Opts, v Vals) Out {
		return b(o, v, f)
	}
}