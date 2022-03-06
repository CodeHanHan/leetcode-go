package goden0402

import (
	"fmt"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	tree := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(tree)
}
