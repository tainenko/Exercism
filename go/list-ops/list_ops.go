package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldl given a function, a list, and initial accumulator, fold (reduce) each item into the accumulator from the left
// using function(accumulator, item)
func (s IntList) Foldl(fn binFunc, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

// Foldr given a function, a list, and an initial accumulator, fold (reduce) each item into the accumulator from the
// right using function(item, accumulator)
func (s IntList) Foldr(fn binFunc, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

// Filter given a predicate and a list, return the list of all items for which predicate(item) is True
func (s IntList) Filter(fn func(int) bool) IntList {
	newLst := IntList{}
	for _, v := range s {
		if fn(v) {
			newLst = append(newLst, v)
		}
	}
	return newLst
}

// Length given a list, return the total number of items within it
func (s IntList) Length() int {
	return len(s)
}

// Map given a function and a list, return the list of the results of applying function(item) on all items
func (s IntList) Map(fn func(int) int) IntList {
	for i := range s {
		s[i] = fn(s[i])
	}
	return s
}

// Reverse given a list, return a list with all the original items, but in reversed order
func (s IntList) Reverse() IntList {
	if len(s) == 0 {
		return s
	}
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

// Append given two lists, add all items in the second list to the end of the first list
func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

// Concat given a series of lists, combine all items in all lists into one flattened list
func (s IntList) Concat(lists []IntList) IntList {
	for _, lst := range lists {
		s = s.Append(lst)
	}
	return s
}
