package leetcode

func isSymmetricWithQueue(root *TreeNode) bool {
	queue := []*TreeNode{}
	queue = append(queue, root.Left, root.Right)
	return checkSymmetric2(queue, 0)
}

func checkSymmetric2(queue []*TreeNode, i int) bool {
	for i != len(queue) {
		l := queue[i]
		i++
		r := queue[i]
		i++
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		queue = append(queue, l.Left)
		queue = append(queue, r.Right)

		queue = append(queue, l.Right)
		queue = append(queue, r.Left)
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	if checkSymmetric(l.Left, r.Right) {
		return checkSymmetric(l.Right, r.Left)
	} else {
		return false
	}
}
