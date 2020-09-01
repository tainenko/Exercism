package letter

// FreqMap is a type of map[rune]int
type FreqMap map[rune]int

// Frequency counts the rune frequency in a string
// and return a FreqMap Object
func Frequency(s string) FreqMap {
	freqMap := FreqMap{}
	for _, c := range s {
		freqMap[c]++
	}
	return freqMap
}

// ConcurrentFrequency receive a string array and call the Frequency function
// with goroutine
func ConcurrentFrequency(StringArray []string) FreqMap {
	ch := make(chan FreqMap, len(StringArray))
	for _, s := range StringArray {
		s := s
		go func() {
			ch <- Frequency(s)
		}()
	}

	freqMap := FreqMap{}
	for range StringArray {
		for k, v := range <-ch {
			freqMap[k] += v
		}
	}
	return freqMap
}
