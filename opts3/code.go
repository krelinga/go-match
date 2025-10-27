package opts3

import (
	"fmt"
	"slices"
)

type Code int

const (
	Err Code = iota
	Yes
	No
)

func (c Code) PanicIfInvalid() {
	switch c {
	case Err, Yes, No:
		// valid
	default:
		panic(fmt.Sprintf("unknown Code value: %d", c))
	}
}

func (c Code) String() string {
	c.PanicIfInvalid()
	switch c {
	case Err:
		return "Err"
	case Yes:
		return "Yes"
	case No:
		return "No"
	default:
		panic("unreachable")
	}
}

func (c Code) Emoji() string {
	c.PanicIfInvalid()
	switch c {
	case Err:
		return "⚠️"
	case Yes:
		return "✅"
	case No:
		return "❌"
	default:
		panic("unreachable")
	}
}

func CodeNot(c Code) Code {
	c.PanicIfInvalid()
	switch c {
	case Err:
		return Err
	case Yes:
		return No
	case No:
		return Yes
	default:
		panic("unreachable")
	}
}

func CodeAnd(codes ...Code) Code {
	if len(codes) == 0 {
		return Err
	}
	for _, c := range codes {
		c.PanicIfInvalid()
		if c == Err {
			return Err
		}
	}
	if slices.Contains(codes, No) {
		return No
	}
	return Yes
}

func CodeOr(codes ...Code) Code {
	if len(codes) == 0 {
		return Err
	}
	for _, c := range codes {
		c.PanicIfInvalid()
		if c == Err {
			return Err
		}
	}
	if slices.Contains(codes, Yes) {
		return Yes
	}
	return No
}
