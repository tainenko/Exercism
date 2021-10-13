package kindergarten

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

// Garden struct
type Garden struct {
	diagram  string
	children []string
}

var plaintMap = map[byte]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

// NewGarden return the garden struct
func NewGarden(diagram string, children []string) (*Garden, error) {
	if strings.HasPrefix(diagram, "\n") == false {
		return nil, errors.New(" window is missing")
	}
	diagram = strings.ReplaceAll(diagram, "\n", "")
	if len(diagram) != len(children)*4 {
		return nil, errors.New(" odd number of cups")
	}
	set := map[string]bool{}
	for _, child := range children {
		if _, ok := set[child]; ok {
			return nil, errors.New("duplicate name")
		}
		set[child] = true
	}
	reg := regexp.MustCompile("[^CGRV]+")
	if reg.MatchString(diagram) {
		return nil, errors.New(" invaid cup codes")
	}
	chi := make([]string, len(children))
	copy(chi, children)
	sort.Strings(chi)
	garden := &Garden{diagram, chi}
	return garden, nil
}

// Plants return plants belong to each student
func (g *Garden) Plants(child string) ([]string, bool) {
	idx := indexOf(g.children, child)
	if idx == -1 {
		return []string{}, false
	}
	return []string{plaintMap[g.diagram[2*idx]], plaintMap[g.diagram[2*idx+1]], plaintMap[g.diagram[2*idx+2*len(g.children)]], plaintMap[g.diagram[2*idx+1+2*len(g.children)]]}, true
}

func indexOf(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
