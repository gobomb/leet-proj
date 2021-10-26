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

	if !isWordInList(beginWord, wordList) {
		wordList = append(wordList, beginWord)
	}

	wordmap := make(map[string]struct{})
	for _, w := range wordList {
		wordmap[w] = struct{}{}
	}

	queue := make([]string, 0)
	visited := make(map[string]struct{})
	queue = append(queue, beginWord)
	visited[beginWord] = struct{}{}
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
