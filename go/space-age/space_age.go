// Package space has a Age method to calculate the earth-years old of planet.
package space

// Planet is a string of planet name
type Planet string

// Age function takes a float64 variable as input and return a Earth-years old in seconds float64  of planet.
func Age(seconds float64, planet Planet) float64 {
	earthYearInSecond := 31557600.0
	plateYearMap := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return seconds / (earthYearInSecond * plateYearMap[planet])
}
