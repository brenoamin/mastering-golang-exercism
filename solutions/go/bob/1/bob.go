// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob implements a lackadaisical teenager.
package bob


import (
    "strings"
    "unicode"
    )

// Hey returns Bob's response to a remark.
func Hey(remark string) string {
	s := strings.TrimSpace(remark)

	if s == "" {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(s, "?")
	isYelling := isYelling(s)

	switch {
	case isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isYelling:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}


func isYelling(s string) bool {
	hasLetter := false

	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if unicode.IsLower(r) {
				return false
			}
		}
	}

	return hasLetter
}
