package offer_35_m_

import (
	"fmt"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	give1 := [][]int{
		{7, -1},
		{13, 0},
		{11, 4},
		{10, 2},
		{1, 0},
	}
	get1 := copyRandom(build(give1))
	fmt.Println(get1)

	give2 := [][]int{
		{1, 1},
		{2, 1},
	}
	get2 := copyRandom(build(give2))
	fmt.Println(get2)

	give3 := [][]int{
		{3, -1},
		{3, 0},
		{3, -1},
	}
	get3 := copyRandom(build(give3))
	fmt.Println(get3)
}

func Test_build(t *testing.T) {
	give := [][]int{
		{7, -1},
		{13, 0},
		{11, 4},
		{10, 2},
		{1, 0},
	}
	get := build(give)
	fmt.Println(get)
}
