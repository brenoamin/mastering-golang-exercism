package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	frequency := map[string]int{}

	for _, v := range normalizedWords(phrase) {
		frequency[v]++
	}
	return frequency
}

func normalizedWords(input string) []string {
	lowered := strings.ToLower(input)

	normalized := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || unicode.IsLetter(r) || isSeparator(r) || r == '\'' {
			return r
		}
		return -1
	}, lowered)

	words := strings.FieldsFunc(normalized, isSeparator)

	result := make([]string, 0, len(words))

	for _, word := range words {
		word = strings.Trim(word, "'")

		if word == "" {
			continue
		}

		result = append(result, word)
	}

	return result
}

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) ||
		strings.ContainsRune(`?!:".;,`, r)
}

