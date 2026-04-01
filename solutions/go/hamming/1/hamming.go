package hamming

import "fmt"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, fmt.Errorf("strands have different length")
    }
    distance := 0
    
	for i := range a {
        if a[i] != b[i] {
            distance++
        }
    }
    return distance, nil
}
