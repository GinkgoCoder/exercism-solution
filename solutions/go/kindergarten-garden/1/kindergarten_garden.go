package kindergarten

import (
	"errors"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

/**
| Plant  | Diagram encoding |
| ------ | ---------------- |
| Grass  | G                |
| Clover | C                |
| Radish | R                |
| Violet | V                |
**/

var plantMap = map[rune]string{
	'V': "violets",
	'G': "grass",
	'R': "radishes",
	'C': "clover",
}

type Garden struct {
	plants   [2][]rune
	children []string
}

func validate(disgram string, children []string) bool {
	lines := strings.Split(disgram, "\n")
	if len(lines) < 3 {
		return false
	}
	if len(lines[1]) != len(children)*2 || len(lines[2]) != len(children)*2 {
		return false
	}
	re := regexp.MustCompile(`[^VGRC\n]`)
	if re.MatchString(disgram) {
		return false
	}
	return true
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	garden := Garden{children: children}
	set := map[string]bool{}
	for _, child := range children {
		if set[child] {
			return nil, errors.New("name duplicated")
		}
		set[child] = true
	}
	index := 0
	if !validate(diagram, children) {
		return nil, errors.New("diagram is not correct")
	}
	for _, line := range strings.Split(diagram, "\n") {
		if line != "" {
			garden.plants[index] = []rune(line)
			index++
		}
	}
	return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	copied := make([]string, len(g.children))
	copy(copied, g.children)
	sort.Strings(copied)
	index := slices.Index(copied, child)
	if index == -1 {
		return []string{}, false
	}
	result := []string{plantMap[g.plants[0][index*2]], plantMap[g.plants[0][index*2+1]], plantMap[g.plants[1][index*2]], plantMap[g.plants[1][index*2+1]]}
	return result, true
}
