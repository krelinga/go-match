package opts3

import "fmt"

type Code int

const (
	Err Code = iota
	Yes
	No
)

func (c Code) String() string {
	switch c {
	case Err:
		return "Err"
	case Yes:
		return "Yes"
	case No:
		return "No"
	default:
		panic(fmt.Sprintf("unknown Code value: %d", c))
	}
}

func (c Code) Emoji() string {
	switch c {
	case Err:
		return "⚠️"
	case Yes:
		return "✅"
	case No:
		return "❌"
	default:
		panic(fmt.Sprintf("unknown Code value: %d", c))
	}
}

func CodeNot(c Code) Code {
	switch c {
	case Err:
		return Err
	case Yes:
		return No
	case No:
		return Yes
	default:
		panic(fmt.Sprintf("unknown Code value: %d", c))
	}
}

func CodeAnd(codes ...Code) Code {
	if len(codes) == 0 {
		return Err
	}
	for _, c := range codes {
		if c == Err {
			return Err
		}
		if c == No {
			return No
		}
	}
	return Yes
}

func CodeOr(codes ...Code) Code {
	if len(codes) == 0 {
		return Err
	}
	for _, c := range codes {
		if c == Err {
			return Err
		}
		if c == Yes {
			return Yes
		}
	}
	return No
}
