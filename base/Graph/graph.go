package graph

import "errors"

// @author: yangchnet
// @create: 2022-03-05 19:16
// @description: 有向图

var (
	ErrInvalidNode = errors.New("非法节点")
	ErrInvalidEdge = errors.New("非法边")
)

// Graph is a directed graph represented by adjacency Matrix
type Graph struct {
	m [][]int // 值为0时表示边不存在
}

func NewGraphByMatrix(matrix [][]int) *Graph {
	return &Graph{
		m: matrix,
	}
}

func NewRawGraph(nodeNum int) *Graph {
	m := make([][]int, nodeNum)
	for i, _ := range m {
		m[i] = make([]int, nodeNum)
	}

	return &Graph{
		m: m,
	}
}

// AddEdge 添加一条从n指向m，权值为v的边
func (g *Graph) AddEdge(n, m int, v int) error {
	if !g.validNode(n) || !g.validNode(m) {
		return ErrInvalidNode
	}
	if v <= 0 {
		return ErrInvalidEdge
	}

	g.m[n][m] = v

	return nil
}

// DelEdge 删除从n指向m的边
func (g *Graph) DelEdge(n, m int) error {
	if !g.validNode(n) || !g.validNode(m) {
		return ErrInvalidNode
	}

	g.m[n][m] = -1

	return nil
}

// Neighbors 获取结点n的所有“邻居”，即一步内可到达n或n可到达的结点
func (g *Graph) Neighbors(n int) ([]int, error) {
	if !g.validNode(n) {
		return nil, ErrInvalidNode
	}

	res := make([]int, 0)
	for i := 0; i < len(g.m); i++ {
		if g.m[n][i] > 0 {
			res = append(res, i)
		}
	}

	for i := 0; i < len(g.m); i++ {
		if i == n { // 防止重复统计指向自身的边
			continue
		}
		if g.m[i][n] > 0 {
			res = append(res, i)
		}
	}

	return res, nil
}

// BFS 广度优先搜索
func (g *Graph) BFS(op func(n int)) {
	visited := make([]bool, len(g.m))
	queue := make(chan int, len(g.m))

	var fn func(n int, g *Graph) = func(n int, g *Graph) {
		if visited[n] {
			return
		}

		queue <- n
		visited[n] = true
		op(n)
		for len(queue) > 0 {
			v := <-queue
			for i := 0; i < len(g.m); i++ {
				if g.m[v][i] > 0 && !visited[i] {
					queue <- i
				}
			}
		}

	}

	for i := 0; i < len(g.m); i++ {
		fn(i, g)
	}

}

// DFS 深度优先搜索
func (g *Graph) DFS(op func(n int)) {
	visited := make([]bool, len(g.m))

	var fn func(n int, g *Graph)
	fn = func(n int, g *Graph) {
		if visited[n] {
			return
		}

		visited[n] = true
		op(n)
		for i := 0; i < len(g.m); i++ {
			if g.m[n][i] > 0 && !visited[i] {
				fn(i, g)
			}
		}
	}

	for i := 0; i < len(g.m); i++ {
		fn(i, g)
	}

}

// HasEdge 检查从n到m是否有边，若有返回边权值，否则返回0
func (g *Graph) HasEdge(n, m int) int {
	if !g.validNode(n) || !g.validNode(m) {
		return -1
	}

	return g.m[n][m]
}

func (g *Graph) validNode(n int) bool {
	if n < 0 || n >= len(g.m) {
		return false
	}

	return true
}
