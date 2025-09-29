package pivot

type Matcher interface {
	Match(got any) (bool, error)
}

type Value interface {
	String() string
	SameTypeAs(other any) bool
}

type ValueComparer interface {
	Value
	Compare(other any) int
}

type ValueEqualer interface {
	Value
	Equal(other any) bool
}
