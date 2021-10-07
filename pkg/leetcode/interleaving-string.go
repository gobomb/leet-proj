package leetcode

/*
	97. Interleaving String
*/

// 1. failed
// Time Limit Exceeded
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if s3 == "" && s1 == s2 && s1 == "" {
		return true
	}

	if s1 != "" && s1[0] == s3[0] && isInterleave(s1[1:], s2, s3[1:]) {
		return true
	}

	return s2 != "" && s2[0] == s3[0] && isInterleave(s1, s2[1:], s3[1:])
}

// 2. use a cache to save the same result
// Runtime: 4 ms, faster than 22.73% of Go online submissions for Interleaving String.
// Memory Usage: 8.7 MB, less than 6.06% of Go online submissions for Interleaving String.

var cacheIsInterleave map[string]bool

func init() {
	cacheIsInterleave = make(map[string]bool)
}

func isInterleaveCache(s1 string, s2 string, s3 string) (b bool) {
	defer func() {
		// log.Printf("%v %v %v %v\n", s1, s2, s3, b)
		if _, ok := cacheIsInterleave[s1+" "+s2+" "+s3]; ok {
			return
		}
		cacheIsInterleave[s1+" "+s2+" "+s3] = b
	}()

	if rs, ok := cacheIsInterleave[s1+" "+s2+" "+s3]; ok {
		return rs
	}

	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if s3 == "" && s1 == s2 && s1 == "" {
		return true
	}

	if s1 != "" && s1[0] == s3[0] && isInterleaveCache(s1[1:], s2, s3[1:]) {
		return true
	}

	b = s2 != "" && s2[0] == s3[0] && isInterleaveCache(s1, s2[1:], s3[1:])
	return b
}

// 3. try let index as args; use 2d array instead of map cache
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Interleaving String.
// Memory Usage: 2.9 MB, less than 28.79% of Go online submissions for Interleaving String.

// 0 not visited
// 1 true
// 2 false
var arrIsInterleave [][]int

func isInterleave2(s1 string, s2 string, s3 string) (b bool) {
	arrIsInterleave = make([][]int, len(s1)+1)
	for i := range arrIsInterleave {
		arrIsInterleave[i] = make([]int, len(s2)+1)
	}

	return isInterleaveCache22(s1, s2, s3, 0, 0)
}

func isInterleaveCache22(s1 string, s2 string, s3 string, i, j int) (b bool) {
	defer func() {
		// log.Printf("%v %v %v %v\n", i, j, arrIsInterleave, b)

		if arrIsInterleave[i][j] == 0 {
			if b {
				arrIsInterleave[i][j] = 1
			} else {
				arrIsInterleave[i][j] = 2
			}
		}
	}()

	if arrIsInterleave[i][j] != 0 {
		return arrIsInterleave[i][j] == 1
	}

	if len(s1[i:])+len(s2[j:]) != len(s3[i+j:]) {
		return false
	}

	if len(s3[i+j:]) == 0 && len(s1[i:]) == len(s2[j:]) && len(s1[i:]) == 0 {
		return true
	}

	if len(s1[i:]) != 0 && s1[i] == s3[i+j] && isInterleaveCache22(s1, s2, s3, i+1, j) {
		return true
	}

	b = len(s2[j:]) != 0 && s2[j] == s3[i+j] && isInterleaveCache22(s1, s2, s3, i, j+1)
	return b
}
