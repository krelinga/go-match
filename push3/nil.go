package push3

type INil interface {
	Nil() bool
}

func MatchNil(in INil, m M[bool]) Result {
	return m(in.Nil())
}
