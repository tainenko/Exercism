package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and stands out his age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, direction string, distance float64, neighbour string) string {
	return fmt.Sprintf("Welcome to my party, %s! You have been assigned to table %X. "+
		"Your table is on the %s, exactly %.1f meters from here. You will be sitting next to %s.",
		name, table, direction, distance, neighbour)
}
