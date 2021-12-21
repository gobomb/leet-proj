package algorithm

import (
	"container/list"
	"justest/pkg/ds"
	"log"
)

// 使用类似广度优先搜索的方法解决赋权图的单源最短路径问题(确定起点的最短路径问题)
// 图怎么表示

/*
	1. 图怎么表示
	2. 测试数据的生成
	3. 算法的实现
	4. 如何表示无穷大

*/

const INF = 0xffffff

func Dijkstra(s int, g *ds.Graph) ([]int, []int) {
	dist := make([]int, g.N())
	path := make([]int, g.N())
	visited := make([]bool, g.N())

	for i := 0; i < g.N(); i++ {
		dist[i] = INF //距离为无穷大
		path[i] = -1  //没有上一个节点
		visited[i] = false
	}

	ind := s

	queue := list.New()
	queue.PushBack(ind)

	path[s] = s
	dist[s] = 0

	// v := s

	// for _, e := range g.Gr[v] {
	// 	if dist[e.To] > dist[e.From]+e.Weight {
	// 		dist[e.To] = dist[e.From] + e.Weight
	// 		path[e.To] = e.From
	// 	}
	// }

	for {
		if queue.Len() == 0 {
			break
		}
		v := queue.Front().Value.(int)
		queue.Remove(queue.Front())

		if visited[v] {
			continue
		}

		visited[v] = true

		for _, e := range g.Gr[v] {
			if dist[e.To] > dist[e.From]+e.Weight {
				dist[e.To] = dist[e.From] + e.Weight
				path[e.To] = e.From
			}
			queue.PushBack(e.To)
		}
	}

	// log.Println(path, dist, visited)

	return path, dist
}

func getPath(path []int, t int) []int {
	pathes := []int{}

	for {
		pathes = append(pathes, t)
		if t == path[t] {
			break
		}
		t = path[t]
	}

	log.Println(pathes)
	rs := []int{}

	for i := len(pathes) - 1; i >= 0; i-- {
		rs = append(rs, pathes[i])
	}

	return rs
}
