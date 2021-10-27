package leetcode

/*
	127. Word Ladder
*/

func isWordInList(s string, wordList []string) bool {
	found := false
	for _, w := range wordList {
		if w == s {
			found = true
		}
	}
	return found
}

/*
Runtime: 152 ms, faster than 65.58% of Go online submissions for Word Ladder.
Memory Usage: 7.5 MB, less than 51.37% of Go online submissions for Word Ladder.
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}

	if !isWordInList(endWord, wordList) {
		return 0
	}

	if beginWord == endWord {
		return 1
	}

	wordmap := make(map[string]struct{})
	for _, w := range wordList {
		wordmap[w] = struct{}{}
	}

	queue := make([]string, 0)
	visited := make(map[string]struct{})
	queue = append(queue, beginWord)
	rs := 1

	for len(queue) != 0 {
		rs++

		size := len(queue)

		for i := 0; i < size; i++ {
			w := queue[0]
			queue = queue[1:]
			// log.Println(w)
			adj := getAdjacents(w, wordmap)
			for _, a := range adj {
				if _, ok := visited[a]; !ok {
					if a == endWord {
						return rs
					}
					queue = append(queue, a)
					visited[a] = struct{}{}
				}
			}
		}

	}
	return 0
}

func getAdjacents(w string, list map[string]struct{}) (adj []string) {
	var l byte
	for l = 'a'; l <= 'z'; l++ {
		for i := 0; i < len(w); i++ {
			if w[i] != l {
				s := w[:i] + string(l) + w[i+1:]
				if _, ok := list[s]; ok {
					adj = append(adj, s)
				}
			}
		}
	}
	return adj
}

/*
Runtime: 24 ms, faster than 99.43% of Go online submissions for Word Ladder.
Memory Usage: 6.6 MB, less than 74.71% of Go online submissions for Word Ladder.
*/
func ladderLengthBiDBFS(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}

	if !isWordInList(endWord, wordList) {
		return 0
	}

	if beginWord == endWord {
		return 1
	}

	wordmap := make(map[string]struct{})
	for _, w := range wordList {
		wordmap[w] = struct{}{}
	}

	beginQueue := []string{beginWord}
	endQueue := []string{endWord}
	visited := make(map[string]struct{})

	rs := 1

	for len(beginQueue) != 0 && len(endQueue) != 0 {
		if len(beginQueue) > len(endQueue) {
			beginQueue, endQueue = endQueue, beginQueue
		}

		rs++

		tmpQueue := []string{}

		for _, w := range beginQueue {
			// log.Println(w)
			adj := getAdjacents(w, wordmap)
			for _, a := range adj {
				if contain(endQueue, a) {
					return rs
				}
				if _, ok := visited[a]; !ok {
					tmpQueue = append(tmpQueue, a)
					visited[a] = struct{}{}
				}
			}
		}
		beginQueue = tmpQueue
	}

	return 0
}

func contain(list []string, s string) bool {
	for _, a := range list {
		if a == s {
			return true
		}
	}
	return false
}
