package goden0401

import "testing"

func Test_findWhetherExistsPath(t *testing.T) {
	findWhetherExistsPath(3, [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}}, 0, 2)
}
