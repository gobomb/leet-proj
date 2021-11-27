package medium

/*
	207. Course Schedule
*/

// https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0207.Course-Schedule/
func canFinish(numCourses int, prerequisites [][]int) bool {
	incount := make([]int, numCourses)
	out := make(map[int][]int)
	queue := make([]int, 0, numCourses)

	for _, v := range prerequisites {
		// 入度的统计
		incount[v[0]]++
		// 出度的记录
		out[v[1]] = append(out[v[1]], v[0])
	}

	for i := 0; i < numCourses; i++ {
		if incount[i] == 0 {
			// 入度为0的节点入队
			queue = append(queue, i)
		}
	}

	// 广度优先搜索
	for i := 0; i != len(queue); i++ {
		c := queue[i]
		for _, v := range out[c] {
			// 删除c到v的边
			incount[v]--
			// 入度为零，入队
			if incount[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	// 输出的顶点数小于网中的顶点数，则有环路
	return len(queue) == numCourses
}
