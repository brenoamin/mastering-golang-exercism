package isbn



func IsValidISBN(isbn string) bool {
	sum := 0
	weight := 10
	count := 0

	for _, r := range isbn {
		switch {
		case r == '-':
			continue

		case r >= '0' && r <= '9':
			sum += int(r-'0') * weight
			weight--
			count++

		case r == 'X':
			if weight != 1 {
				return false
			}
			sum += 10
			weight--
			count++

		default:
			return false
		}
	}

	return count == 10 && sum%11 == 0
}