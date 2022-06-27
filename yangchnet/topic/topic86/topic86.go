package topic86

import (
	"math"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func partition(head *ListNode, x int) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	_head := &ListNode{Val: math.MinInt32, Next: head}

// 	lm, xptr := _head, head // lm 是xptr之前最后一个值小于x的节点
// 	for xptr != nil && xptr.Val != x {
// 		if xptr.Val < x {
// 			lm = xptr
// 		}
// 		xptr = xptr.Next
// 	}

// 	if xptr == nil {
// 		return head
// 	}

// 	// 将所有小于x的节点移动到x之前
// 	for pre, rptr := xptr, xptr.Next; rptr != nil; {
// 		if rptr.Val < x { // rptr为小于x的节点
// 			// 临时记录要断链位置的节点
// 			lmNext := lm.Next
// 			rptrNext := rptr.Next

// 			// 重新链接
// 			lm.Next = rptr
// 			rptr.Next = lmNext
// 			pre.Next = rptrNext

// 			// 更新指针
// 			lm = rptr
// 			rptr = pre.Next
// 			continue
// 		}

// 		pre = rptr
// 		rptr = rptr.Next
// 	}

// 	_lm := lm
// 	for pre, p := _head, _head.Next; p != nil && p.Val != x; {
// 		if p == _lm { // 所有大于等于x的节点都被移动到小于x的节点之后了
// 			break
// 		}
// 		if p.Val > x {
// 			if p == lm.Next {
// 				break
// 			}

// 			lmNext := lm.Next
// 			pNext := p.Next

// 			lm.Next = p
// 			p.Next = lmNext
// 			pre.Next = pNext

// 			lm = p
// 			p = pre.Next
// 			continue
// 		}

// 		pre = p
// 		p = p.Next
// 	}

// 	return _head.Next
// }

func partition(head *ListNode, x int) *ListNode {
	minList := &ListNode{math.MinInt64, nil}
	maxList := &ListNode{math.MaxInt64, nil}
	_head := &ListNode{math.MinInt64, head}
	minPtr, maxPtr, p, pre := minList, maxList, head, _head

	for p != nil {
		if p.Val < x {
			minPtr.Next = p
			pre.Next = p.Next
			p.Next = nil

			minPtr = p
			p = pre.Next
		} else {
			maxPtr.Next = p
			pre.Next = p.Next
			p.Next = nil

			maxPtr = p
			p = pre.Next
		}
	}

	minPtr.Next = maxList.Next

	return minList.Next
}
