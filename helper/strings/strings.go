package strings

import "strings"

func ContainsFold(s, m string) bool {
	return strings.Contains(
		strings.ToLower(s),
		strings.ToLower(m),
	)
}
