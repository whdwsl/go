// Package weather : weather management.
// can forecast the weather.
package weather

//CurrentCondition: the weather condition.
var CurrentCondition string

//CurrentLocation: the location.
var CurrentLocation string

// Forecast : forecast the weather condition in the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
