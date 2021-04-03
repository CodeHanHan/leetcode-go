package main

const TAILINDEX = 0

type Queue struct { //右边是尾，左边是头
	val  []*TreeNode
	tail int //永远指向0
	head int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newQueue() *Queue {
	return &Queue{
		val:  []*TreeNode{},
		tail: TAILINDEX,
		head: 0,
	}
}

func newTree(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (queue *Queue) EnQueue(val *TreeNode) { //入队，放在最右边
	queue.val = append(queue.val, val)
	queue.head++
}

func (queue *Queue) isEmpty() bool {
	return queue.head == queue.tail
}

func (queue *Queue) DeQueue() *TreeNode { //出队，从最左边出去
	if queue.isEmpty() {
		panic("empty queue")
	}
	res := queue.val[0]
	queue.val = queue.val[1:]
	queue.head--
	return res
}

func levelOrder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return []int{}
	}
	queue := newQueue()
	queue.EnQueue(root)
	for !queue.isEmpty() {
		cur := queue.DeQueue()
		res = append(res, cur.Val)
		if cur.Left != nil {
			queue.EnQueue(cur.Left)
		}
		if cur.Right != nil {
			queue.EnQueue(cur.Right)
		}
	}
	return res
}

//Line 25: Char 9: cannot use []*TreeNode literal (type []*TreeNode) as type []*leetcode.TreeNode in field value (solution.go)
//Line 40: Char 23: cannot use val (type *TreeNode) as type *leetcode.TreeNode in append (solution.go)
//Line 55: Char 5: cannot use res (type *leetcode.TreeNode) as type *TreeNode in return argument (solution.go)
//Line 103: Char 37: cannot use param_1 (type *leetcode.TreeNode) as type *TreeNode in argument to __helper__ (solution.go)
