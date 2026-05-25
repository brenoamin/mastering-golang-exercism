package series

func All(n int, s string) []string {
	if n > len(s) || n <=0 {
		return  []string{}
	}
	var result []string
	runes := []rune(s)

	for l, r := 0, n-1; r+1 <= len(s); l, r = l+1, r+1 {
        currentSubstring := ""
		for _, number := range runes[l : r+1] {
			currentSubstring += string(number)
		}
		result = append(result, currentSubstring)
	}

	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}

	return s[:n], true
}
