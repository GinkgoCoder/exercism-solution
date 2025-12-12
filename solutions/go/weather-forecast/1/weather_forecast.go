// Package weather provides functionality for weather forecasting.
// It maintains the current weather condition and location information.
package weather

var (
    // CurrentCondition stores the current weather condition.
	CurrentCondition string
    // CurrentLocation stores the current location being monitored.
	CurrentLocation  string
)

// Forecast sets the current location and weather condition, 
// and returns a formatted string describing the weather forecast for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
