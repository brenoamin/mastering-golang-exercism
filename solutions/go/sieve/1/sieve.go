package sieve

import (
	"slices"
)

func Sieve(limit int) []int {
	sieveCriteria := map[int]bool{}

	for i := 2; i <= limit; i++ {
		sieveCriteria[i] = true
	}

	for k := range sieveCriteria {
		if sieveCriteria[k] {
			for i := k * 2; i <= limit; i = i + k {
				sieveCriteria[i] = false
			}
		}
	}

	var un []int

	for k, v := range sieveCriteria {
		if v {
			un = append(un, k)
		}
	}
    slices.Sort(un)
    
	return un

}
