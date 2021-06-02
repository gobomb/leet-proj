package leetcode

func generateTrees(n int) []*TreeNode {
	return genBSTrees(1, n)
}

func genBSTrees(s, e int) []*TreeNode {
	var rs []*TreeNode
	if s > e {
		rs = append(rs, nil)
		return rs
	}
	for i := s; i <= e; i++ {
		left := genBSTrees(s, i-1)
		right := genBSTrees(i+1, e)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				}
				rs = append(rs, root)
			}
		}
	}
	return rs
}

// Catalan number
// https://zh.wikipedia.org/wiki/%E5%8D%A1%E5%A1%94%E5%85%B0%E6%95%B0
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}
