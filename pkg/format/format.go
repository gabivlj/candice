package format

import "strings"

func StringWithTabs(s string, startingTabs int) string {
	builder := strings.Builder{}
	tabs := startingTabs
	prev := 'a'
	for _, c := range s {
		if c == '}' {
			tabs--
		}
		if c == '{' {
			tabs++
		}
		if prev == '\n' {
			builder.WriteString(strings.Repeat("\t", tabs))
		}
		builder.WriteRune(c)
		prev = c
	}
	return builder.String()
}
