package machine

import (
	"mcsim/data"
	"mcsim/state"
	"mcsim/util"
)

type Machine struct {
	states []*state.State
}

func (m *Machine) Walk(idx int, size int) state.Path {
	lst := make(state.Path, size)
	m.states[idx].Walk(lst, size, 0)
	return lst
}

func New() Machine {
	ans := Machine{
		states: make([]*state.State, data.MatrixSize),
	}
	for i := 0; i < data.MatrixSize; i++ {
		ans.states[i] = state.New(util.ID(i))
	}
	for i := 0; i < data.MatrixSize; i++ {
		for j := 0; j < data.MatrixSize; j++ {
			ans.states[i].AddInput(ans.states[j], data.Matrix[i*data.MatrixSize+j])
		}
		ans.states[i].Initialize()
	}
	return ans
}
