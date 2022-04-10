package disjointset

// 并查集
// 并查集用来解决一些元素分组的问题。它管理一系列不相交的集合，并支持两种操作：
// 1.合并（Union）：把两个不相交的集合合并为一个集合。
// 2. 查询（Find）：查询两个元素是否在同一个集合中。

// Reference: https://zhuanlan.zhihu.com/p/93647900

type DisjointSet struct {
	arr  []int // arr 存储每个节点的父节点
	rank []int // rank存储每个节点的秩，记录每个根节点对应的树的深度（如果不是根节点，其rank相当于以它作为根节点的子树的深度）
}

// NewDisjointSet create a new set
func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		arr:  []int{-1},
		rank: []int{-1},
	}
}

// Init init set
func (s *DisjointSet) Init(n int) {
	for i := 1; i <= n; i++ {
		s.arr = append(s.arr, i)
		s.rank = append(s.rank, 1)
	}
}

// Find find root parent for node x
func (s *DisjointSet) Find(x int) int {
	if s.arr[x] == x {
		return x
	}

	s.arr[x] = s.Find(s.arr[x]) // 路径压缩，将父节点直接指向root
	return s.arr[x]
}

// Merge merge node i and node j
func (s *DisjointSet) Merge(i, j int) {
	x, y := s.Find(i), s.Find(j) // 先找到两个根节点
	if x <= y {
		s.arr[x] = y
	} else {
		s.arr[y] = x
	}

	// 如果深度相同且根节点不同，则新的根节点的深度+1
	if s.rank[x] == s.rank[y] && x != y {
		s.rank[y]++
	}
}

// 例题:
// 题目背景： 若某个家族人员过于庞大，要判断两个是否是亲戚，确实还很不容易，现在给出某个亲戚关系图，求任意给出的两个人是否具有亲戚关系。
// 题目描述: 规定：x和y是亲戚，y和z是亲戚，那么x和z也是亲戚。如果x,y是亲戚，那么x的亲戚都是y的亲戚，y的亲戚也都是x的亲戚。
// Input: n int, p [][]int, q [][]int
// n个人, p为亲戚关系, q对人，求他们是否具有亲戚关系
// Outpu: []bool
/// 输入的q对人，他们的亲戚关系

func findRelative(n int, p [][]int, q [][]int) []bool {
	s := NewDisjointSet()
	s.Init(n)

	for _, _p := range p {
		s.Merge(_p[0], _p[1])
	}

	ret := make([]bool, 0)
	for _, _q := range q {
		ret = append(ret, s.Find(_q[0]) == s.Find(_q[1]))
	}

	return ret
}
