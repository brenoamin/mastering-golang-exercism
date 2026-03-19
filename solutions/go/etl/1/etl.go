package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    out := make(map[string]int, len(in))
    
	for index, value := range in {
        for _, letter := range value {
            lower := strings.ToLower(letter)
        	out[lower] = index
        }
    }
    
    return out
}
