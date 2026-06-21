package romannumerals

import "fmt"

type OrderedMap struct {
	keys []int
	data map[int]string
}

func NewOrderedMap(keys []int, data map[int]string) *OrderedMap {
	return &OrderedMap{
		keys: keys,
		data: data,
	}
}

func (om *OrderedMap) Get(key int) (string, bool) {
	value, ok := om.data[key]
	return value, ok
}
func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", fmt.Errorf("%v is out of range", input)
	}

	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	mapKeyValue := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	om := NewOrderedMap(keys, mapKeyValue)

	result := input
	roman := ""
	i := 0

	for {
		if result <= 0 {
			break
		}
		if result-keys[i] >= 0 {
			number, ok := om.Get(keys[i])
			if ok {
				roman += number
				result = result - keys[i]
			}
		} else {
			i++
		}

	}
	return roman, nil
}
