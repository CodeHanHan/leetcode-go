package topic958

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := make(chan *TreeNode, 100)
	queue <- root

	var miss bool = false
	for len(queue) > 0 {
		cur := <-queue
		if miss && (cur.Left != nil || cur.Right != nil) { // 当前节点之前有缺失子节点的节点
			return false
		}

		if cur.Left == nil && cur.Right != nil { // 左空右非空
			return false
		}

		if cur.Left == nil || cur.Right == nil { // 存在缺失，设置标志位
			miss = true
		}

		if cur.Left != nil {
			queue <- cur.Left
		}

		if cur.Right != nil {
			queue <- cur.Right
		}
	}

	return true

}
