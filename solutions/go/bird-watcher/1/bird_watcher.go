package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.

func TotalBirdCount(birdsPerDay []int) int {
    total := 0
	for _, b := range birdsPerDay {
        total+=b
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    from := (week - 1) * 7
    to := from + 7

    if from >= len(birdsPerDay) {
        return 0
    }
    if to > len(birdsPerDay) {
        to = len(birdsPerDay)
    }
    total := 0
    for _, b := range birdsPerDay[from:to] {
        total += b
    }
    return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i, _ := range birdsPerDay{
        if (i%2 == 0){
            birdsPerDay[i]++
        }
    }
    return birdsPerDay
}
