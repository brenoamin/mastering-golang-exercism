package bottlesong

import (
    "fmt"
    "strings"
)

var numbers = []string{
	"No",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
}

func Recite(startBottles, takeDown int) []string {
	var verses []string

	for i := startBottles; i > startBottles-takeDown; i-- {
		current := numbers[i]
		next := strings.ToLower(numbers[i-1])

		bottle := "bottles"
		if i == 1 {
			bottle = "bottle"
		}

		nextBottle := "bottles"
		if i-1 == 1 {
			nextBottle = "bottle"
		}

		verses = append(verses,
			fmt.Sprintf("%s green %s hanging on the wall,", current, bottle),
			fmt.Sprintf("%s green %s hanging on the wall,", current, bottle),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green %s hanging on the wall.", next, nextBottle),
		)

		if i > startBottles-takeDown+1 {
			verses = append(verses, "")
		}
	}

	return verses
}