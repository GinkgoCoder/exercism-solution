package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var names = map[string]bool{}

const maxNameNum = 26 * 26 * 10 * 10 * 10

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(names) == maxNameNum {
		return "", errors.New("No more names can be generated")
	}
	var name string
	for name == "" || names[name] {
		name = fmt.Sprintf("%c%c%d%d%d", rand.Intn(26)+'A', rand.Intn(26)+'A', rand.Intn(10), rand.Intn(10), rand.Intn(10))
	}
	r.name = name
	names[name] = true
	return name, nil
}

func (r *Robot) Reset() {
	r.name = ""
	delete(names, r.name)
}
