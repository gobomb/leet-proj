package leetcode

/*
126. Word Ladder II
*/

/*
Runtime: 4 ms, faster than 100.00% of Go online submissions for Word Ladder II.
Memory Usage: 3.3 MB, less than 82.35% of Go online submissions for Word Ladder II.

https://leetcode-cn.com/problems/word-ladder-ii/solution/yan-du-you-xian-bian-li-shuang-xiang-yan-du-you--2/
*/

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := &[][]string{}

	if len(wordList) == 0 {
		return [][]string{}
	}

	if !isWordInList(endWord, wordList) {
		return [][]string{}
	}

	if beginWord == endWord {
		return [][]string{{beginWord}}
	}

	// beginword 预先入队，要先从剩余节点去除
	wordList = remove(wordList, beginWord)

	// 存剩余的节点
	wordset := make(map[string]struct{})
	for _, w := range wordList {
		wordset[w] = struct{}{}
	}

	// 记录每个节点所在的层数
	steps := make(map[string]int)

	steps[beginWord] = 0

	// 存图数据
	// 记录每个节点可以从哪些节点扩展而来
	from := make(map[string][]string)

	if bfsFindLadders(beginWord, endWord, wordset, steps, from) {
		path := []string{endWord}
		dfsFindLadders(beginWord, endWord, from, res, path)
	}
	return *res
}

// bfs 构建图（from）
func bfsFindLadders(beginWord, endWord string, wordset map[string]struct{}, steps map[string]int, from map[string][]string) bool {
	step := 0
	found := false

	queue := []string{beginWord}

	for len(queue) != 0 {
		step++

		for _, currWord := range queue {
			queue = queue[1:]

			for j, l := range currWord {
				for c := 'a'; c <= 'z'; c++ {
					if l != c {
						// 扩展节点
						nextWord := currWord[:j] + string(c) + currWord[j+1:]
						if level, ok := steps[nextWord]; ok && level == step {
							// nextWord 已经不在 wordset 中了（已经有至少一个父节点了），但还可能有父节点
							// 一个节点有多个父节点
							from[nextWord] = append(from[nextWord], currWord)
						}

						// 已经被处理过了
						if _, ok := wordset[nextWord]; !ok {
							continue
						}

						delete(wordset, nextWord)

						// 入队
						queue = append(queue, nextWord)

						// 记录父节点（第一次）
						from[nextWord] = append(from[nextWord], currWord)

						steps[nextWord] = step

						if nextWord == endWord {
							found = true
						}
					}
				}
			}
		}
		if found {
			break
		}
	}
	return found
}

// dfs 查找所有路径（回溯）
func dfsFindLadders(beginWord, cur string, from map[string][]string, res *[][]string, path []string) {
	if cur == beginWord {
		*res = append(*res, path)
		return
	}

	for _, pre := range from[cur] {
		path = append([]string{pre}, path...)
		dfsFindLadders(beginWord, pre, from, res, path)
		path = path[1:]
	}
}

func remove(list []string, s string) []string {
	for i, a := range list {
		if a == s {
			list[i] = list[len(list)-1]
			list = list[:len(list)-1]
			break
		}
	}
	return list
}
