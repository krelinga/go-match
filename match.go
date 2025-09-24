package match

func Match[T any](m Matcher[T], input T) bool {
	r := &boolReporter{}
	m.Match(input, r)
	return r.Matched()
}

type boolReporter struct {
	mismatch bool
}

func (r *boolReporter) Report(_ string) {
	r.mismatch = true
}

func (r *boolReporter) Child(_ Namer) Reporter {
	return r
}

func (r *boolReporter) Matched() bool {
	return !r.mismatch
}
