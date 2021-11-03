package sublist

import (
	"reflect"
)

// Relation between two list
type Relation string

const (
	equal     Relation = "equal"
	sublist   Relation = "sublist"
	superlist Relation = "superlist"
	unequal   Relation = "unequal"
)

// Sublist return relation between two list
func Sublist(listOne, listTwo []int) Relation {
	if len(listOne) == len(listTwo) {
		if reflect.DeepEqual(listOne, listTwo) {
			return equal
		}
		return unequal
	}
	if len(listOne) < len(listTwo) && isSub(listOne, listTwo) {
		return sublist
	}
	if len(listOne) > len(listTwo) && isSub(listTwo, listOne) {
		return superlist
	}
	return unequal
}
func isSub(listOne, listTwo []int) bool {
	if len(listOne) == 0 {
		return true
	}
	for i := 0; i <= len(listTwo)-len(listOne); i++ {
		if listOne[0] == listTwo[i] && reflect.DeepEqual(listOne, listTwo[i:i+len(listOne)]) {
			return true
		}
	}
	return false
}
