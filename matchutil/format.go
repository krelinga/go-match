package matchutil

import (
	"fmt"
	"reflect"
	"strings"
)

func Emoji(matched bool) string {
	if matched {
		return "✅"
	}
	return "❌"
}

func TypeName(x any) string {
	return reflect.TypeOf(x).String()
}

func FormatWith[T any](t T, format func(t T) string) string {
	if format != nil {
		return format(t)
	}
	return fmt.Sprintf("%v", t)
}

func Indent(s string, level int) string {
	prefix := strings.Repeat(" ", level*3)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

func Explain(matched bool, matcherName string, details ...string) string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "%s %s", Emoji(matched), matcherName)
	if len(details) > 0 {
		sb.WriteString(":")
		for _, detail := range details {
			sb.WriteString("\n")
			sb.WriteString(Indent(strings.TrimSpace(detail), 1))
		}
	}
	return sb.String()
}

func ActualVsExpected(actual, expected string) string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "Actual:   %s\n", actual)
	fmt.Fprintf(sb, "Expected: %s", expected)
	return sb.String()
}

func Describe(in any) string {
	if s, ok := in.(fmt.Stringer); ok {
		return s.String()
	}
	return fmt.Sprintf("%#v", in)
}
