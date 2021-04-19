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
		r.name = getNewName()
	}
	return r.name, nil
}

func getNewName() string {
	var name string
	for {
		name = getStringWithCharset(2, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") + getStringWithCharset(3, "0123456789")
		if !nameMap[name] {
			break
		}
	}
	nameMap[name] = true
	return name
}

//Reset function would set the Robot name as empty
func (r *Robot) Reset() {
	r.name = getNewName()
}
