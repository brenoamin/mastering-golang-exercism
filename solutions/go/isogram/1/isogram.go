package isogram

import (
	"strings"
	"unicode"
)


func IsIsogram(word string) bool {
	normalized := normalize(word)
	seen := make(map[int]bool)
	for _, l := range normalized {
		_, ok := seen[int(l)]
		if ok {
			return false
		} else {
			seen[int(l)] = true
		}
	}
	return true
}

func normalize(s string) string {
	lower := strings.ToLower(s)
	
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			return r
		} else {
			return -1
		}
	}, lower)
}