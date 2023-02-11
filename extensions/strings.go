package extensions

import "strings"

func ContainsCaseInsensitive(a string, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}
