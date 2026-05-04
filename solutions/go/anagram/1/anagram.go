package anagram

import (
	"strings"
)


func Detect(subject string, candidates []string) []string {
	subjectBase := map[rune]int{}

	for _, l := range strings.ToLower(subject) {
		subjectBase[l]++
	}
	var output []string

	for _, word := range candidates {
		base := map[rune]int{}
		for _, l := range strings.ToLower(word) {
			base[l]++
		}
		if mapsEqual(subjectBase, base) && strings.ToLower(word) != strings.ToLower(subject) {
			output = append(output, word)
		}
	}

	return output
}

func mapsEqual(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
