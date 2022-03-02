package goden0207

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getIntersectionNode(t *testing.T) {
	l0 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: l0,
		},
	}
	l2 := &ListNode{
		Val:  1,
		Next: l0,
	}

	require.Equal(t, 3, getIntersectionNode(l1, l2).Val)
}
