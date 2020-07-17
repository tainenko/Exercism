// Package leap has a IsLeapYear to judge the year is leap or not
package leap

// IsLeapYear function takes a integer as input and return true if the input is a leap year.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%400 == 0 {
			return true
		} else if year%100 == 0 {
			return false
		}
		return true
	}
	return false
}
