package match

import "slices"

func Explain[T any](m Matcher[T], input T) Results {
	results := Results{}
	m.Match(input, &resultReporter{out: &results, namer: m})
	return results
}

type Results []Result

func (r Results) Matched() bool {
	return len(r) == 0
}

type Result struct {
	Path    Path
	Message string
}

type Path []string

type resultReporter struct {
	out    *Results
	namer  Namer
	parent *resultReporter
}

func (r *resultReporter) Report(message string) {
	path := Path{r.namer.Name()}
	for p := r.parent; p != nil; p = p.parent {
		path = append(path, p.namer.Name())
	}
	slices.Reverse(path)
	*r.out = append(*r.out, Result{Path: path, Message: message})
}

func (r *resultReporter) Child(namer Namer) Reporter {
	return &resultReporter{
		out:    r.out,
		namer:  namer,
		parent: r,
	}
}
