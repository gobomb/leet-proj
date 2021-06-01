package leetcode

func generateTrees(n int) []*TreeNode {
	return bfsGenTrees(1, n)
}

func bfsGenTrees(s, e int) []*TreeNode {
	var rs []*TreeNode
	if s > e {
		rs = append(rs, nil)
		return rs
	}
	for i := s; i <= e; i++ {

		left := bfsGenTrees(s, i-1)
		right := bfsGenTrees(i+1, e)
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
