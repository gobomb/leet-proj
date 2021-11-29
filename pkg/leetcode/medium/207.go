package medium

/*
	207. Course Schedule
*/

// https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0207.Course-Schedule/
func canFinish(numCourses int, prerequisites [][]int) bool {
	queue := findOrder(numCourses, prerequisites)

	// 输出的顶点数小于网中的顶点数，则有环路
	return len(queue) == numCourses
}
