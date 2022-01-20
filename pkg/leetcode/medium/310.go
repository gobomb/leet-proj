package medium

/*
	310. Minimum Height Trees
*/

func findMinHeightTrees(n int, edges [][]int) []int {
	var rs []int
	type set struct {
		degree int
		nodes  []int
	}

	degree := make(map[int]*set)

	updateDegree := func(a, b int) {
		d := degree[a]

		if d == nil {
			degree[a] = &set{}
			d = degree[a]
		}

		d.degree++
		d.nodes = append(d.nodes, b)
	}

	for _, edge := range edges {
		updateDegree(edge[0], edge[1])
		updateDegree(edge[1], edge[0])
	}

	if n == 1 {
		degree[0] = &set{degree: 0}
	}

	for {
		if len(degree) == 1 || len(degree) == 2 || len(degree) == 0 {
			for node := range degree {
				rs = append(rs, node)
			}
			return rs
		}

		toDel := []int{}

		for k := range degree {
			if degree[k].degree == 1 {
				toDel = append(toDel, k)
			}
		}

		for _, delNode := range toDel {
			d := degree[delNode]
			delete(degree, delNode)

			edge := d.nodes[0]

			degree[edge].degree--
			nodes := degree[edge].nodes

			for i := 0; i < len(nodes); i++ {
				if nodes[i] == delNode {
					nodes[i] = nodes[len(nodes)-1]
					degree[edge].nodes = nodes[:len(nodes)-1]
					break
				}
			}
		}
	}
}
