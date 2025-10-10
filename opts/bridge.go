package opts

type Bridge func(Options, Values, Func) (bool, error)

func (b Bridge) And(m Func) Func {
	return func(opts Options, vals Values) (bool, error) {
		return b(opts, vals, m)
	}
}