package machine

import (
	"fmt"
	"mcsim/data"
	"mcsim/state"
	"mcsim/util"
	"strings"
)

type AnalysisResult struct {
	state.Path
	Probabilities []float32
}

type AnalysisResultSet struct {
	results []AnalysisResult
	size    int
}

func (res AnalysisResultSet) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("PATH LENGTH %d\n", res.size))
	for startState, result := range res.results {
		sb.WriteString(fmt.Sprintf("\tStart point %d:\n", startState+1))
		sb.WriteString(fmt.Sprintf("\t\tPath = %s\n", result.Path))
		sb.WriteString(fmt.Sprintf("\t\tProbabilities:\n\t\t\t%6s %6s %6s\n", "State", "Need", "Got"))
		for i, got := range result.Probabilities {
			expected := data.ExpectedProbabilities[i]
			sb.WriteString(fmt.Sprintf("\t\t\t%6d %6s %6s\n", i+1, util.FormatFloat(expected, 6), util.FormatFloat(got, 6)))
		}
	}
	return sb.String()
}

func (m *Machine) Analyze(size int) AnalysisResultSet {
	res := make([]AnalysisResult, len(m.states))
	for i := range m.states {
		res[i].Probabilities = make([]float32, len(m.states))
		res[i].Path = m.Walk(i, size)
		for _, step := range res[i].Path {
			res[i].Probabilities[step.ID]++
		}
		for j := range res[i].Probabilities {
			res[i].Probabilities[j] /= float32(size)
		}
	}
	return AnalysisResultSet{res, size}
}
