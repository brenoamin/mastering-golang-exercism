package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	result := 0
	for i, v := range colors[0:2] {
		if i == 0 {
			result = colorToNumber(v) * 10
		} else {
			result += colorToNumber(v)
		}
	}
	return result
}

func colorToNumber(color string) int {
	mapColorToNumber := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	return mapColorToNumber[color]
}
