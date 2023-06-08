package data

import "mcsim/util"

const MatrixSize = 5

var Matrix = [...]util.Probability{
	1, 3, 0, 3, 3,
	2, 1, 0, 4, 3,
	3, 2, 5, 0, 0,
	3, 5, 0, 1, 1,
	4, 1, 0, 2, 3,
}

var ExpectedProbabilities = [...]float32{
	0.25, 0.25, 0, 0.25, 0.25,
}
