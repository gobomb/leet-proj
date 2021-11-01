package leetcode

func findMode(root *TreeNode) []int {
	var rs []int
	m := make(map[int]int)
	max := 0

	computeMode(root, m)

	for _, v := range m {
		if v > max {
			max = v
		}
	}

	for k, v := range m {
		if v == max {
			rs = append(rs, k)
		}
	}
	return rs
}

func computeMode(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}

	m[root.Val]++

	computeMode(root.Left, m)
	computeMode(root.Right, m)
}
