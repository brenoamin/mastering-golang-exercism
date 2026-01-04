package atbash

import (
    "strings"
    "unicode"
)

const CIPHER = "zyxwvutsrqponmlkjihgfedcba"

func Atbash(s string) string {
  var b strings.Builder
    count := 0

	for _, r := range s {
		switch {
		case unicode.IsLetter(r):
			r = unicode.ToLower(r)
			if count == 5 {
				b.WriteByte(' ')
				count = 0
			}
			b.WriteByte(CIPHER[r-'a'])
			count++

		case unicode.IsDigit(r):
			if count == 5 {
				b.WriteByte(' ')
				count = 0
			}
			b.WriteRune(r)
			count++
		}
	}
	return b.String()
}