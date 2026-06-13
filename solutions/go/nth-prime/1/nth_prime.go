package nthprime

import (
    "fmt"
    "math"
)
// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("prime cannot be computed for numbers below 1")
	}

	var primeNumbers []int
	i := 2
	for {
		sqrt := int(math.Sqrt(float64(i)))
		divisors := 0
		for c := 2; c <= sqrt; c++ {
			if i%c == 0 {
				divisors++
			}
		}
		if divisors == 0 {
			primeNumbers = append(primeNumbers, i)
		}
		if len(primeNumbers) == n {
			return primeNumbers[n-1], nil
		}
		divisors = 0
		i++
	}
}
