package matchfmt

import (
	"fmt"
	"strings"
)

// Emoji returns a visual indicator emoji based on the match result.
// It returns "✅" for true (matched) and "❌" for false (not matched).
func Emoji(matched bool) string {
	if matched {
		return "✅"
	}
	return "❌"
}

// IndentBy indents each line of the input string s by the specified level.
// Each level adds 3 spaces of indentation. A level of 0 returns the original string.
func IndentBy(s string, level int) string {
	prefix := strings.Repeat(" ", level*3)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

// Indent indents each line of the input string s by one level (3 spaces).
// This is equivalent to calling IndentBy(s, 1).
func Indent(s string) string {
	return IndentBy(s, 1)
}

// Explain formats a matcher explanation with an emoji, matcher name, and optional details.
// The matched parameter determines the emoji (✅ or ❌), matcherName is displayed after the emoji,
// and details are indented on separate lines if provided.
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

// ActualVsExpected formats actual and expected values for comparison display.
// It returns a formatted string showing "Expected: <expected>" on the first line
// and "Actual: <actual>" on the second line.
func ActualVsExpected(actual, expected string) string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "Expected: %s\n", expected)
	fmt.Fprintf(sb, "Actual:   %s", actual)
	return sb.String()
}
