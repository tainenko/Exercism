package flatten

// Flatten accepts an arbitrarily-deep nested list-like structure and returns a flattened structure without any nil/null values.
func Flatten(nested interface{}) []interface{} {
	flatten := []interface{}{}
	nestedSlice, ok := nested.([]interface{})
	if !ok {
		return flatten
	}
	for _, ele := range nestedSlice {
		switch ele.(type) {
		case []interface{}:
			flatten = append(flatten, Flatten(ele)...)
		case nil:
			continue
		default:
			flatten = append(flatten, ele)
		}
	}
	return flatten

}
