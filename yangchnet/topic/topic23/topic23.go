package topic23

import "math"

func mergeKLists(lists []*ListNode) *ListNode {
	buckets := make(map[int][]*ListNode)

	var min, max int = math.MaxInt32, math.MinInt32
	for _, list := range lists {
		var p *ListNode = list
		for ; p != nil; p = p.Next {
			if min > p.Val {
				min = p.Val
			}
			if max < p.Val {
				max = p.Val
			}
			buckets[p.Val] = append(buckets[p.Val], p)
		}
	}

	var head *ListNode = &ListNode{}
	var cur *ListNode = head
	for i := min; i <= max; i++ {
		if list, ok := buckets[i]; ok {
			for _, node := range list {
				cur.Next = node
				cur = cur.Next
			}
		}
	}

	return head.Next
}
