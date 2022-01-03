package easy

/*
	997. Find the Town Judge
*/

func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}

	trusted := make(map[int]struct {
		from int
		to   int
	})

	for _, p := range trust {
		f := trusted[p[1]]
		f.from++
		trusted[p[1]] = f

		j := trusted[p[0]]
		j.to++
		trusted[p[0]] = j
	}

	for k, v := range trusted {
		if v.from == n-1 && v.to == 0 {
			return k
		}
	}

	return -1
}

func findJudge1(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}

	trusted := make([]int, n+1)

	for _, p := range trust {
		trusted[p[1]]++
		trusted[p[0]]--
	}

	for k, v := range trusted {
		if v == n-1 {
			return k
		}
	}

	return -1
}
