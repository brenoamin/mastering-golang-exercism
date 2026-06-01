package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	nid := normalize(id)
	if len(nid) <= 1 || !containsOnlyNumbers(nid) {
		return false
	}
	nid = removeSpaces(nid)
	result := 0
	shouldDouble := false
	for i := len(nid) - 1; i >= 0; i-- {
		currentNumber, err := strconv.Atoi(string(nid[i]))
		if err != nil {
			panic("Couldn't convert string to int")
		}
		if shouldDouble {
			currentNumber *= 2
			if currentNumber > 9 {
				currentNumber -= 9
			}
		}
		result += currentNumber
		shouldDouble = !shouldDouble
	}
	return result%10 == 0
}

func normalize(input string) string {
	return strings.TrimSpace(input)
}

func containsOnlyNumbers(input string) bool {
	for _, n := range input {
		if (n >= '0' && n <= '9') || n == ' ' {
			continue
		} else {
			return false
		}
	}
	return true
}

func removeSpaces(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		} else {
			return -1
		}
	}, input)
}