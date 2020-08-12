package clock

import "fmt"

// Clock has one attribute ,minute.
type Clock struct {
	minute int
}

// New would return a Clock instance
func New(hour, minute int) Clock {
	minute = (hour*60 + minute) % 1440
	if minute < 0 {
		minute += 1440
	}
	return Clock{minute}
}

// String is a stringer of Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

// Add provide a Clock adding method
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minute+minutes)
}

// Subtract provide a Clock subtracting method
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minute-minutes)
}
