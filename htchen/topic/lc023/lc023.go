package lc023

func mergeKLists1(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoLists(ans, lists[i])
	}
	return ans
}

func mergeKLists2(list []*ListNode) *ListNode {
	return merge(list, 0, len(list)-1)
}

func merge(list []*ListNode, l int, r int) *ListNode {
	if l == r {
		return list[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoLists(merge(list, l, mid), merge(list, mid+1, r))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	prehead := &ListNode{}
	pre := prehead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}

	if list1 == nil {
		pre.Next = list2
	}
	if list2 == nil {
		pre.Next = list1
	}

	return prehead.Next
}
