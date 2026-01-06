package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
        return 0, errors.New("The number must be greather or equal than 1")
    }
	steps := 0
    for n != 1 {
        if n%2 == 0 {
            n/= 2
        } else {
            n = 1+n*3
        }
        steps++
    }
    return steps, nil
}
