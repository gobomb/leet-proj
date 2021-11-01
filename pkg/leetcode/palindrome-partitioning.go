package leetcode

/*
	131. Palindrome Partitioning
*/

var ppDp [][]int

func partitionPalindrome(s string) [][]string {
	res := [][]string{}
	ps := []string{}
	ppDp = make([][]int, len(s)+1)
	for i := range ppDp {
		ppDp[i] = make([]int, len(s)+1)
	}
	partitionPalindrome2(s, ps, 0, &res)
	return res
}

func partitionPalindrome2(s string, ps []string, start int, res *[][]string) {
	if len(s) == start {
		*res = append(*res, append([]string{}, ps...))
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome3(s, start, i+1) {
			ps = append(ps, s[start:i+1])
			partitionPalindrome2(s, ps, i+1, res)
			ps = ps[:len(ps)-1]
		}
	}
}

func isPalindrome3(ss string, m, n int) bool {
	if ppDp[m][n] != 0 {
		return ppDp[m][n] == 1
	}

	s := ss[m:n]
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			ppDp[m][n] = -1
			return false
		}
		i++
		j--
	}

	ppDp[m][n] = 1
	return true
}
