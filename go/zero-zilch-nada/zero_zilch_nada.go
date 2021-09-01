package zero

// EmptyBool returns an empty (zero value) bool
func EmptyBool() bool {
	return false
}

// EmptyInt returns an empty (zero value) int
func EmptyInt() int {
	return 0
}

// EmptyString returns an empty (zero value) string
func EmptyString() string {
	return ""
}

// EmptyFunc returns an empty (zero value) func
func EmptyFunc() func() {
	var ret func()
	return ret
}

// EmptyPointer returns an empty (zero value) pointer
func EmptyPointer() *int {
	var p *int
	return p
}

// EmptyMap returns an empty (zero value) map
func EmptyMap() map[int]int {
	var ret map[int]int
	return ret
}

// EmptySlice returns an empty (zero value) slice
func EmptySlice() []int {
	var ret []int
	return ret
}

// EmptyChannel returns an empty (zero value) channel
func EmptyChannel() chan int {
	var ch chan int
	return ch
}

// EmptyInterface returns an empty (zero value) interface
func EmptyInterface() interface{} {
	var ret interface{}
	return ret
}
