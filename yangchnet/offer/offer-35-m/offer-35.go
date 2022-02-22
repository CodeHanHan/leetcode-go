package offer_35_m_

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandom(head *Node) *Node {
	if head == nil {
		return nil
	}
	var m map[*Node]int = make(map[*Node]int) // 存储节点对应的位序
	var s []*Node                             // 存储新节点
	p := head

	for i := 0; ; i++ {
		if p == nil {
			s = append(s, nil)
			m[nil] = i
			break
		}
		n := newNode(p.Val)
		s = append(s, n)
		m[p] = i
		p = p.Next
	}

	p = head
	for i := 0; i < len(s)-1; i++ {

		if i == len(s)-2 { // 最后一个非nil节点
			s[i].Next = nil
		} else {
			s[i].Next = s[i+1]
		}
		s[i].Random = s[m[p.Random]] // 设置其Random指针
		p = p.Next
	}

	return s[0]
}

func newNode(Val int) *Node {
	return &Node{
		Val:    Val,
		Next:   nil,
		Random: nil,
	}
}

func build(give [][]int) *Node {
	get := make([]*Node, 0)
	for _, v := range give { // 构造节点
		get = append(get, newNode(v[0]))
	}
	for i, v := range give { // 链接Random
		if v[1] == -1 {
			get[i].Random = nil
			continue
		}
		get[i].Random = get[v[1]]
	}
	for i, v := range get { // 链接Next
		if i == len(get)-1 {
			v.Next = nil
			continue
		}
		v.Next = get[i+1]
	}
	return get[0]
}

func (n *Node) show() [][]int {
	m := make(map[*Node]int)
	res := make([][]int, 0)
	p := n
	for i := 0; ; i++ { // 设置map
		if p == nil {
			m[nil] = -1
			break
		}
		m[p] = i
		p = p.Next
	}

	p = n
	for i := 0; ; i++ {
		if p == nil {
			break
		}
		r := make([]int, 2)
		r[0] = p.Val
		r[1] = m[p.Random]
		res = append(res, r)
		p = p.Next
	}
	return res
}

func (n *Node) String() string {
	m := make(map[*Node]int)
	res := make([][]int, 0)
	p := n
	for i := 0; ; i++ { // 设置map
		if p == nil {
			m[nil] = -1
			break
		}
		m[p] = i
		p = p.Next
	}

	p = n
	for i := 0; ; i++ {
		if p == nil {
			break
		}
		r := make([]int, 2)
		r[0] = p.Val
		r[1] = m[p.Random]
		res = append(res, r)
		p = p.Next
	}
	return fmt.Sprint(res)
}
