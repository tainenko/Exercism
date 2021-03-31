package robotname

import "math/rand"

type Robot struct {
	name string
}

func randSeq(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randDigit(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += (string)(rune(rand.Intn(10) + 48))
	}
	return s
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		prefix := randSeq(2)
		suffix := randDigit(3)
		r.name = prefix + suffix
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
