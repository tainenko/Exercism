package robotname

import (
	"math/rand"
	"time"
)

// Robot struct has a name attribute
type Robot struct {
	name string
}

var nameMap = make(map[string]bool)
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func getStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Name function give the Robot instance an unused random name.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name := "xxxx"
		nameMap[name] = true
		for nameMap[name] {
			name = getStringWithCharset(2, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") + getStringWithCharset(3, "0123456789")
		}
		nameMap[name] = true
		r.name = name
	}
	return r.name, nil
}

//Reset function would set the Robot name as empty
func (r *Robot) Reset() {
	r.name = ""
}
