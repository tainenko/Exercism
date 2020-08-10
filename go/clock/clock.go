package clock

import "fmt"

// Clock has two integer attributes: hour and minute.
type Clock struct {
	hour   int
	minute int
}

// New function is a Clock constructor
func New(hour, minute int) Clock {
	c := Clock{hour, minute}
	c.balance()
	return c
}

// String is a stringer of Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add provide a clock adding method
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	c.balance()
	return c
}

// Subtract provide a clock subtracting method
func (c Clock) Subtract(minutes int) Clock {
	c.minute -= minutes
	c.balance()
	return c
}

func (c *Clock) balance() {
	total := (60*c.hour + c.minute) % 1440
	if total < 0 {
		total += 1440
	}
	c.hour = total / 60
	c.minute = total % 60
}
