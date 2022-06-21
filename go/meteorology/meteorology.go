package meteorology

import "fmt"

// TemperatureUnit is the unit of temperature
type TemperatureUnit int

// the unit of temperature will be either `Celsius` or `Fahrenheit`
const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (t TemperatureUnit) String() string {
	if t == Celsius {
		return "°C"
	}
	return "°F"
}

// Temperature values consist of an integer and a temperature unit
type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit.String())
}

// SpeedUnit is the unit of wind speed
type SpeedUnit int

// the unit of wind speed will be either `KmPerHour` or `MilesPerHour`
const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (s SpeedUnit) String() string {
	if s == KmPerHour {
		return "km/h"
	}
	return "mph"
}

// Speed consist of an integer and a speed unit
type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit.String())
}

// Data specifies location, temperature, wind direction, wind speed and humidity
type Data struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m Data) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", m.location, m.temperature.String(), m.windDirection, m.windSpeed.String(), m.humidity)
}
