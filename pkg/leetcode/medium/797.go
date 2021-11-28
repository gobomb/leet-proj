package medium

/*
	797. All Paths From Source to Target
*/

func allPathsSourceTarget(graph [][]int) [][]int {
	res := []int{0}
	ress := [][]int{}

	allPathsSourceTarget1(graph, res, 0, &ress)

	return ress
}

func allPathsSourceTarget1(graph [][]int, res []int, ind int, ress *[][]int) {
	if res[len(res)-1] == len(graph)-1 {
		cp := make([]int, len(res))
		copy(cp, res)
		*ress = append(*ress, cp)
		return
	}

	for _, g := range graph[ind] {
		res = append(res, g)
		allPathsSourceTarget1(graph, res, g, ress)
		res = res[:len(res)-1]
	}
}
