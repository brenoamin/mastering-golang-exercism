package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, avgTimePerLayer int) int {
    if avgTimePerLayer == 0 {
        avgTimePerLayer = 2
    }
    return len(layers)*avgTimePerLayer
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
    var noodles int
    var sauce int
    
    for _,v := range layers {
        switch v {
            case "noodles":
            	noodles++
            case "sauce":
            	sauce++
        }
    }
    
    return noodles*50, float64(sauce)*0.2 
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendRecipe, myRecipe []string) {
    lastIngredientFriendRecipe := friendRecipe[len(friendRecipe)-1]
    myRecipe[len(myRecipe)-1] = lastIngredientFriendRecipe
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(amountForTwo []float64, portions int) []float64 {
    amountNeeded := make([]float64, len(amountForTwo))
    
    timesFactor := float64(portions)/2
    for i, v := range amountForTwo {
        amountNeeded[i] = v * timesFactor
    }

    return amountNeeded
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
