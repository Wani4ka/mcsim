package state

import (
	"fmt"
	"mcsim/data"
	"strings"
)

type Path []*State

func (p Path) String() string {
	sz := len(p)
	if sz <= data.MaxSize {
		return fmt.Sprintln([]*State(p))
	}
	p1 := []*State(p[:data.MaxSize/2])
	p2 := []*State(p[sz-(data.MaxSize+1)/2:])
	return fmt.Sprintf("[%s ... %s]", strings.Trim(fmt.Sprint(p1), "[]"), strings.Trim(fmt.Sprint(p2), "[]"))
}
