package meetup

import "time"

// WeekSchedule indicates the day of week begins
type WeekSchedule int

const (
	// First week begins on 7
	First WeekSchedule = 1
	// Second week begins on 8
	Second WeekSchedule = 8
	// Third week begins on 15
	Third WeekSchedule = 15
	// Fourth week begins on 22
	Fourth WeekSchedule = 22
	// Teenth begins on thirteen
	Teenth WeekSchedule = 13
	// Last week begins on 6 days before next month
	Last WeekSchedule = -6
)

// Day calculates the date of meetups.
func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	if week == Last {
		month++
	}
	t := time.Date(year, month, int(week), 12, 0, 0, 0, time.UTC)
	return t.Day() + (int(weekday-t.Weekday())+7)%7
}
