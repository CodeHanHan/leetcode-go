package base

import (
	"log"
	"plugin"

	list "github.com/CodeHanHan/leetcode-go/lichang/base/LinkList"
)

func LoadLinkList(filename string) (
	ListNode list.ListNode,
	BuildListWithHead func([]int) *list.ListNode,
	BuildListWithNoHead func([]int) *list.ListNode,
) {
	p, err := plugin.Open(filename)
	if err != nil {
		log.Fatalf("cannot load plugin %v", filename)
	}

	xListNode, err := p.Lookup("ListNode")
	if err != nil {
		log.Fatalf("cannot find ListNode %v", filename)
	}
	ListNode = xListNode.(list.ListNode)

	xBuildListWithHead, err := p.Lookup("BuildListWithHead")
	if err != nil {
		log.Fatalf("cannot find BuildListWithHead %v", filename)
	}
	BuildListWithHead = xBuildListWithHead.(func([]int) *list.ListNode)

	xBuildListWithNoHead, err := p.Lookup("BuildListWithNoHead")
	if err != nil {
		log.Fatalf("cannot find BuildListWithHead %v", filename)
	}
	BuildListWithNoHead = xBuildListWithNoHead.(func([]int) *list.ListNode)
	return
}
