// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
    "strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	words := splitBySpaceOrHyphens(s)
	acronym := ""

	for _, word := range words {
		for i, l := range word {
			if i == 0 {
				acronym += strings.ToUpper(string(l))
			} else {
				break
			}
		}
	}
	return acronym
}

func splitBySpaceOrHyphens(s string) []string {
	result := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	return result
}