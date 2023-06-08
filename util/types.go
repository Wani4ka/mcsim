package util

import "strconv"

type Probability uint
type ID int

func (id ID) String() string {
	return strconv.Itoa(int(id + 1))
}
