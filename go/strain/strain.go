package strain

// Ints is a collection of integer
type Ints []int

// Keep create a new Ints which satisfies the condition
func (i *Ints) Keep(f func(int) bool) Ints {
	var kept Ints
	for _, v := range *i {
		if f(v) {
			kept = append(kept, v)
		}
	}
	return kept
}

// Discard create a new Ints which doesn't satisfy the condition
func (i *Ints) Discard(f func(int) bool) Ints {
	var kept Ints
	for _, v := range *i {
		if !f(v) {
			kept = append(kept, v)
		}
	}
	return kept
}

// Lists is a collection of integer array
type Lists [][]int

// Keep create a new Lists which satisfies the condition
func (l *Lists) Keep(f func([]int) bool) Lists {
	var kept Lists
	for _, v := range *l {
		if f(v) {
			kept = append(kept, v)
		}
	}
	return kept
}

// Strings is a collection of string
type Strings []string

// Keep create a new Strings which satisfies the condition
func (s *Strings) Keep(f func(string) bool) Strings {
	var kept Strings
	for _, v := range *s {
		if f(v) {
			kept = append(kept, v)
		}
	}
	return kept
}

// Discard create a new Strings that doesn't satisfy the condition
func (s *Strings) Discard(f func(string) bool) Strings {
	var kept Strings
	for _, v := range *s {
		if !f(v) {
			kept = append(kept, v)
		}
	}
	return kept
}
