package accum

import "strings"

func Accum(s string) string {
	mumbles := make([]string, len(s))
	for i, r := range s {
		mumbles[i] = strings.ToUpper(string(r)) + strings.ToLower(strings.Repeat(string(r), i))
	}
	return strings.Join(mumbles, "-")
}
