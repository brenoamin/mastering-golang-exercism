package lineup

import (
    "strconv"
)

func Format(name string, number int) string {
	
	message := name + ", you are the " + strconv.Itoa(number) + suffixDecisor(number) + " customer we serve today. Thank you!"
	
	return message
}

func suffixDecisor(number int) string {
	strNumber := strconv.Itoa(number)
	suffix := ""
	endingNumber := string(strNumber[len(strNumber)-1])
	penultimateNumber := ""

	if len(strNumber) > 1 {
		penultimateNumber = string(strNumber[len(strNumber)-2])
	}

	if endingNumber != "1" && endingNumber != "2" && endingNumber != "3" {
		suffix = "th"
	} else {
		if endingNumber == "1" {
			if penultimateNumber == "1" {
				suffix = "th"
			} else {
				suffix = "st"
			}
		}
		if endingNumber == "2" {
			if penultimateNumber == "1" {
				suffix = "th"
			} else {
				suffix = "nd"
			}
		}
		if endingNumber == "3" {
			if penultimateNumber == "1" {
				suffix = "th"
			} else {
				suffix = "rd"
			}
		}
	}

	return suffix
}
