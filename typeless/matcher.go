package typeless

import (
	"fmt"
	"runtime"
	"strings"
)

type Matched bool

func (m Matched) Emoji() string {
	if m {
		return "✅"
	}
	return "❌"
}

type Explanation string

type Matcher interface {
	Match(got any) (Matched, Explanation, error)
}

type FmtFunc func(got any) (string, error)

func DefaultFmt(got any) (string, error) {
	return fmt.Sprintf("%#v", got), nil
}

var (
	ErrType  = fmt.Errorf("type error")
	ErrValue = fmt.Errorf("value error")
)

func Error(base error, message string) error {
	stack := []string{}
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stack = append(stack, fmt.Sprintf("%s:%d", file, line))
	}
	stackStr := strings.Join(stack, "\n")
	return fmt.Errorf("%w: %s\nStack trace:\n%s", base, message, stackStr)
}

type FuncMatcher func(got any) (Matched, Explanation, error)

func (f FuncMatcher) Match(got any) (Matched, Explanation, error) {
	return f(got)
}
