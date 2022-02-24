package offer07

import (
	"fmt"
	"testing"
)

func Test_buildTree(t *testing.T) {
	preorder1 := []int{3, 9, 20, 15, 7}
	inorder1 := []int{9, 3, 15, 20, 7}
	root1 := buildTree(preorder1, inorder1)
	fmt.Println(root1)

	preorder2 := []int{-1}
	inorder2 := []int{-1}
	root2 := buildTree(preorder2, inorder2)
	fmt.Println(root2)
}
