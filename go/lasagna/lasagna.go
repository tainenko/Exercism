package lasagna

// OvenTime function that does not take any parameters and returns how many minutes the lasagna should be in the oven
func OvenTime() int {
	return 40
}

// RemainingOvenTime function that takes the actual minutes the lasagna has been in the oven as a parameter and returns
// how many minutes the lasagna still has to remain in the oven, based on the expected oven time in minutes from the
// previous task.
func RemainingOvenTime(number int) int {
	return OvenTime() - number
}

// PreparationTime function that takes the number of layers you added to the lasagna as a parameter and returns how many
// minutes you spent preparing the lasagna, assuming each layer takes you 2 minutes to prepare.
func PreparationTime(layer int) int {
	preparingTimeInMinute := 2
	return layer * preparingTimeInMinute
}

// ElapsedTime function that takes two parameters: the first parameter is the number of layers you added to the lasagna,
// and the second parameter is the number of minutes the lasagna has been in the oven.
func ElapsedTime(layer, minute int) int {
	return PreparationTime(layer) + minute
}
