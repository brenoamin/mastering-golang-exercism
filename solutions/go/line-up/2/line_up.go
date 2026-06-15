package lineup

import (
    "strconv"
)

func Format(name string, number int) string {
	
	message := name + ", you are the " + strconv.Itoa(number) + suffixDecisor(number) + " customer we serve today. Thank you!"
	
	return message
}

func suffixDecisor(number int) string {
	lastOne := number % 10
	lastTwo := number % 100

	if lastTwo >= 11 && lastTwo <= 13 {
		return "th"
	}

	switch lastOne {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}
