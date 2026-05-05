package primefactors

func Factors(n int64) []int64 {
	factor := 2
	var result []int64

	for {
		if n == 1 {
			break
		}
		if n%int64(factor) == 0 {
			n = n / int64(factor)
			result = append(result, int64(factor))
		} else {
			factor++
		}
	}
	return result
}
