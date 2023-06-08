package state

import (
	"fmt"
	"github.com/mroth/weightedrand/v2"
	"log"
	"mcsim/util"
)

type Probability uint

type State struct {
	ID      util.ID
	choices []weightedrand.Choice[*State, util.Probability]
	chooser *weightedrand.Chooser[*State, util.Probability]
}

func (st *State) String() string {
	return fmt.Sprint(st.ID)
}

func New(id util.ID) *State {
	return &State{
		ID: id,
	}
}

func (st *State) AddInput(target *State, p util.Probability) {
	st.choices = append(st.choices, weightedrand.NewChoice(target, p))
}

func (st *State) Initialize() {
	var err error
	st.chooser, err = weightedrand.NewChooser(st.choices...)
	if err != nil {
		log.Fatalf("Couldn't initialize chooser for state %s: %s", st.ID, err)
	}
}

func (st *State) Walk(lst []*State, size, idx int) {
	if idx >= size {
		return
	}
	lst[idx] = st
	st.chooser.Pick().Walk(lst, size, idx+1)
}
