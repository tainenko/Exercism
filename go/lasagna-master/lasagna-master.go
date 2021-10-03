package lasagna

// PreparationTime return the estimate for the total preparation time based on the number of layers as an int.
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

const quantityOfNoodles = 50
const quantityOfSauce = 0.2

// Quantities determines the quantity of noodles and sauce needed to make your meal.
func Quantities(layers []string) (int, float64) {
	var quantitiesOfNoodles int
	var quantitiesOfSauce float64
	for _, layer := range layers {
		if layer == "noodles" {
			quantitiesOfNoodles += quantityOfNoodles
		}
		if layer == "sauce" {
			quantitiesOfSauce += quantityOfSauce
		}
	}
	return quantitiesOfNoodles, quantitiesOfSauce
}

// AddSecretIngredient unction should generate a new slice and add the last item from your
// friends list to the end of your list.
func AddSecretIngredient(friendsList, mylist []string) []string {
	return append(mylist, friendsList[len(friendsList)-1])
}

const portionOfSlice = 2.0

// ScaleRecipe return a slice of float64s with the amounts needed for the desired number of portions.
func ScaleRecipe(quantities []float64, portion int) []float64 {
	var amounts []float64
	scale := float64(portion) / portionOfSlice
	for _, quantity := range quantities {
		amounts = append(amounts, quantity*scale)
	}
	return amounts
}
