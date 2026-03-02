package thefarm

import (
    "errors"
    "fmt"
)


type InvalidCowsError struct {
    message string
    count int
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.count, e.message)
}

// TODO: define the 'DivideFood' function

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
    amount, err := fc.FodderAmount(cows)
    if err != nil {
        return 0, err
    }
    factor, errFc := fc.FatteningFactor()
    if errFc != nil {
        return 0, errFc
    }
    
    return (amount*factor)/float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
    if cows > 0 {
        amount, err := DivideFood(fc, cows)
        if err != nil {
            return 0, err
        }
        return amount, nil
    }
    return 0, errors.New("invalid number of cows")
}


// TODO: define the 'ValidateNumberOfCows' function

func ValidateNumberOfCows(cows int) error {
    if cows < 0 {
        return &InvalidCowsError{
            message: "there are no negative cows",
            count: cows,
        }
    }
    if cows == 0 {
        return &InvalidCowsError{
            message: "no cows don't need food",
            count: cows,
        }
    }
    return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
