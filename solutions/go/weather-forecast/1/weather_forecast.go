// Package weather This package is used to forcast weather.
package weather

var (
	// CurrentCondition represents the weather status.
	CurrentCondition string
	// CurrentLocation represents the location where we want the forcast.
	CurrentLocation string
)

// Forecast returns the status of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
