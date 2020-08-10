package clock

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) *Clock {
	c := new(Clock)
	c.hour = hour
	c.minute = minute
	return c
}

func (c *Clock) String() string {
	return ""
}

func (c *Clock) Add(minutes int) *Clock {
	c.minute += minutes
	if c.minute >= 60 {
		c.minute -= 60
		c.hour += 1
	}
	return c
}

func (c *Clock) Subtract(minutes int) *Clock {
	c.minute -= minutes
	if c.minute < 0 {
		c.minute += 60
		c.hour -= 1
	}
	return c
}
