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

func (queue *Queue) Clear() {
	//queue = newQueue()   //这个写法为啥不管用
	//解释
	/*
	   因为
	*/

	queue.val = queue.val[0:0]
	queue.head = 0
}

func newTree(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (queue *Queue) length() int {
	return queue.head
}

func (queue *Queue) EnQueue(val *TreeNode) { //入队，放在最右边
	queue.val = append(queue.val, val)
	queue.head++
}

func (queue *Queue) isEmpty() bool { //若队列空返回true
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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	q1 := newQueue()
	q2 := newQueue()
	flag := 0
	q1.EnQueue(root)
	for !q1.isEmpty() || !q2.isEmpty() {
		r := []int{}
		//var q *Queue
		if flag == 0 {
			for i := 0; i < q1.length(); i++ {
				r = append(r, q1.val[i].Val)
				if q1.val[i].Left != nil {
					q2.EnQueue(q1.val[i].Left)
				}
				if q1.val[i].Right != nil {
					q2.EnQueue(q1.val[i].Right)
				}
			}
			q1.Clear()
		} else {
			for i := 0; i < q2.length(); i++ {
				r = append(r, q2.val[i].Val)
				if q2.val[i].Left != nil {
					q1.EnQueue(q2.val[i].Left)
				}
				if q2.val[i].Right != nil {
					q1.EnQueue(q2.val[i].Right)
				}

			}
			q2.Clear()
		}
		res = append(res, r)
		flag = (flag + 1) % 2
	}
	return res
}
