package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	set := map[int]bool{}

	for _, v := range divisors {
		if v == 0 {
			continue
		}
		count := 1
		for {
			if v*count < limit {
				set[v*count] = true
			} else {
				break
			}
			count++
		}
	}
	sum := 0

	for n := range set {
		sum += n
	}
	return sum
}
