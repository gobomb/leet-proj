package ds

import (
	"fmt"
)

//
type Graph struct {
	n        int  // 节点数
	m        int  // 边数量
	directed bool //有向 or 无向图
	Gr       [][]*E
}

type E struct {
	From   int //a节点
	To     int //b节点
	Weight int //权值
}

func NewGraph(n int, d bool) *Graph {
	return &Graph{
		n:        n,
		directed: d,
		Gr:       make([][]*E, n),
	}
}

func (g *Graph) N() int {
	return g.n
}

func (g *Graph) AddEdge(v1, v2 int, weight int) {
	if v1 < 0 || v1 >= g.n || v2 < 0 || v2 >= g.n {
		return
	}

	g.Gr[v1] = append(g.Gr[v1], &E{v1, v2, weight})

	if v1 != v2 && !g.directed {
		// 无向图，添加反方向
		g.Gr[v2] = append(g.Gr[v2], &E{v2, v1, weight})
	}

	g.m++
}

func (g *Graph) String() string {
	s := "\n"

	for i, from := range g.Gr {
		s += fmt.Sprintf("%v: ", i)
		for _, to := range from {
			s += fmt.Sprintf("%v - %v	", to.To, to.Weight)
		}
		s += fmt.Sprintln()
	}

	return s
}
