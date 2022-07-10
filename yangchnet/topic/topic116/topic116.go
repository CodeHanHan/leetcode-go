package topic116

// . "github.com/CodeHanHan/leetcode-go/base/Tree"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	q := []*Node{root}

	for len(q) != 0 {
		q[len(q)-1].Next = nil
		for i := len(q) - 2; i >= 0; i-- {
			q[i].Next = q[i+1]
		}

		p := []*Node{}
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				p = append(p, q[i].Left)
			}
			if q[i].Right != nil {
				p = append(p, q[i].Right)
			}
		}
		q = p
	}

	return root
}
