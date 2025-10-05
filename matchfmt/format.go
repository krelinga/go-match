package matchfmt

import (
	"fmt"
	"strings"
)

func Emoji(matched bool) string {
	if matched {
		return "✅"
	}
	return "❌"
}

func IndentBy(s string, level int) string {
	prefix := strings.Repeat(" ", level*3)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

func Indent(s string) string {
	return IndentBy(s, 1)
}

func Explain(matched bool, matcherName string, details ...string) string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "%s %s", Emoji(matched), matcherName)
	if len(details) > 0 {
		sb.WriteString(":")
		for _, detail := range details {
			sb.WriteString("\n")
			sb.WriteString(Indent(detail))
		}
	}
	return sb.String()
}

func ActualVsExpected(actual, expected string) string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "Expected: %s\n", expected)
	fmt.Fprintf(sb, "Actual:   %s", actual)
	return sb.String()
}

func DefaultFormat(t any) string {
	return fmt.Sprintf("%#v", t)
}
