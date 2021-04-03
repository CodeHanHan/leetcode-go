package base

import (
	"fmt"
	"testing"
)

func TestBuildList(t *testing.T) {
	l := BuildList([]int{1, 2, 3, 4, 5})
	fmt.Println(l)
}
