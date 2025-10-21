//Package weather forecast the current weather conditions of various cities in Goblinocus.
package weather

var (
    //CurrentCondition is a variable responsible for storing the weather condition.
	CurrentCondition string
    //CurrentLocation is a variable responsible for storing the weather condition.
	CurrentLocation  string
)

//Forecast is a function that returns the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
