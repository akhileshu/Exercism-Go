// Package weather is about weather forecast.
package weather



var (
    // CurrentCondition is string.
	CurrentCondition string
    // CurrentLocation is string.
	CurrentLocation  string
)

// Forecast function gives Forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
