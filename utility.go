package UntisAPI

import (
	"strings"
)

/*
TODO what do these functions do
*/
func splitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return removeEmptyStrings(strings.FieldsFunc(s, splitter)...)
}

func removeEmptyStrings(strings ...string) []string {
	for i := len(strings) - 1; i >= 0; i-- {
		if strings[i] == "" {
			strings = append(strings[:i], strings[i+1:]...)
		}
	}
	return strings
}
