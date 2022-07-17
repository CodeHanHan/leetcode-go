package lc116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		p := []*Node{}
		for i := 0; i < len(q); i++ {
			node := q[i]
			if i+1 < len(q) {
				node.Next = q[i+1]
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return root
}
