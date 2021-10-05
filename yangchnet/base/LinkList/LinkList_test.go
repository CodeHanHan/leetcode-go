package list

import (
	"fmt"
	"testing"
)

func TestBuildList(t *testing.T) {
	l := BuildListWithHead([]int{1, 2, 3, 4, 5})
	l1 := BuildListWithNoHead([]int{1, 2, 3, 4, 5})

	fmt.Println(l)
	fmt.Println(l1)

	fmt.Println(l.Slice())
	fmt.Println(l1.Slice())
}
