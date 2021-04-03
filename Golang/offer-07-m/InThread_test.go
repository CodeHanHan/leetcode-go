package offer_07_m_

import "testing"

func TestThreadTreeNode_inThread(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTreeThread(preorder, inorder)
	root.inThread()
}
