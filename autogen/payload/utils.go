package payload

import "strings"

func httpRequestEscape(content string) string {
	var builder strings.Builder
	for _, r := range content {
		switch r {
		case '\b':
			builder.WriteString(`\b`)
		case '\f':
			builder.WriteString(`\f`)
		case '\n':
			builder.WriteString(`\n`)
		case '\r':
			builder.WriteString(`\r`)
		case '\t':
			builder.WriteString(`\t`)
		case '"':
			builder.WriteString(`\"`)
		case '\'':
			builder.WriteString(`\'`)
		case '\\':
			builder.WriteString(`\\`)
		default:
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
