package allergies

var items = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats"}

// Allergies returns the full list of allergies by given allergies score
func Allergies(allergies uint) []string {
	var res []string
	for i := 0; i < len(items); i++ {
		if allergies&(1<<i) != 0 {
			res = append(res, items[i])
		}
	}
	return res
}

// AllergicTo shows whether the given score is allergic to specific allergen
func AllergicTo(allergies uint, allergen string) bool {
	for _, a := range Allergies(allergies) {
		if a == allergen {
			return true
		}
	}
	return false
}
