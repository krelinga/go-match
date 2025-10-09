package push3

type ILen interface {
	Len() int
}

func MatchLen(in ILen, m M[int]) Result {
	return m(in.Len())
}
