package backtrack

import "testing"

func Test_EightQueen(t *testing.T) {
	c := &CheckerBoard{
		matrix: [8]int{},
	}

	c.cal8queen(0)
}
