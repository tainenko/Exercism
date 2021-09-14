// Package weather provides a Log function to record the weather condition of a specific location
package weather

var (
	// CurrentCondition represents the weather condition
	CurrentCondition string
	// CurrentLocation represents the location
	CurrentLocation string
)

// Log is the record of weather condition in a specific location
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
