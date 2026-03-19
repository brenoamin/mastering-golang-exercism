package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    total := 0
    for _, letters := range in {
        total += len(letters)
    }
    out := make(map[string]int, total)

	for score, letters := range in {
        for _, letter := range letters {
            lower := strings.ToLower(letter)
        	out[lower] = score
        }
    }
    
    return out
}
