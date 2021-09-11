package offer_07_m_

import (
	"fmt"
	"testing"
)

func Test_BuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := BuildTree(preorder, inorder)
	fmt.Println(root)
}
