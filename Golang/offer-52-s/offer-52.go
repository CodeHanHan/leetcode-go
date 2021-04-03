package offer_52

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return q
}

func buildList(nums []int) *ListNode {
	a := build(nums, 0)
	return a
}

func build(nums []int, index int) *ListNode { // 不要头节点
	if index == len(nums) {
		return nil
	}
	return &ListNode{
		Val:  nums[index],
		Next: build(nums, index+1),
	}
}
