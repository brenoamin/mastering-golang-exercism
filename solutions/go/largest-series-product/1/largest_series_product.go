package largestseriesproduct

import (
	"fmt"
	"strconv"
)


func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span <= 0 || span > len(digits) || !isASCIIOnlyNumbers(digits) {
		return 0, fmt.Errorf("invalid span provided")
	}
	runes := []rune(digits)
	var largestSeries int64
	for l, r := 0, span-1; r+1 <= len(digits); l, r = l+1, r+1 {
		partialResult := int64(1)
		for _, number := range runes[l : r+1] {
			num, err := strconv.Atoi(string(number))
			if err != nil {
				panic("str to int conversion returned err")
			}
			partialResult *= int64(num)
		}
		if partialResult >= largestSeries {
			largestSeries = partialResult
		}
	}
	return largestSeries, nil
}

func isASCIIOnlyNumbers(s string) bool {
	for _, n := range s {
		if n < '0' || n > '9' {
			return false
		}
	}
	return true
}
