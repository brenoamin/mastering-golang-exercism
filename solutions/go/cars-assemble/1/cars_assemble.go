package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    var percentageSuccessRate = float64(successRate/100)
	return float64(productionRate)*percentageSuccessRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateWorkingCarsPerHour(productionRate, successRate)/60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(totalCars int) uint {
	groupsOfTen := totalCars / 10
	remainingCars := totalCars % 10

	totalCost := groupsOfTen*95000 + remainingCars*10000
	return uint(totalCost)
}
