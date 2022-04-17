package split

import "strings"

type Stringer interface {
	String() string
}

func Split[T Stringer](stringers []T, separator string) string {
	builder := strings.Builder{}

	for i, t := range stringers {
		if i >= 1 {
			builder.WriteString(separator)
		}

		builder.WriteString(t.String())
	}

	return builder.String()
}
