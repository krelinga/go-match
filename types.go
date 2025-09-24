package match

type Namer interface {
	Name() string
}

type Matcher[T any] interface {
	Namer
	Match(input T, reporter Reporter)
}

type Formatter[T any] interface {
	Format(input T) string
}

type Reporter interface {
	Report(message string)
	Child(namer Namer) Reporter
}