package leetcode

import (
	"sort"
)

func recoverTree(root *TreeNode) {
	var prev, first, second *TreeNode
	_, first, second = findFaultBST(root, prev, first, second)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

func findFaultBST(root, prev, first, second *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if root == nil {
		return prev, first, second
	}
	prev, first, second = findFaultBST(root.Left, prev, first, second)
	if prev != nil && prev.Val > root.Val {
		if first == nil {
			first = prev
		}
		second = root
	}
	prev = root
	prev, first, second = findFaultBST(root.Right, prev, first, second)
	return prev, first, second
}

func recoverTree2(root *TreeNode) {
	rs := []int{}
	recursInorderTraversal(root, &rs)
	sort.Ints(rs)
	i := 0
	fixTree(root, rs, &i)
}

func fixTree(root *TreeNode, rs []int, i *int) {
	if root == nil {
		return
	}

	fixTree(root.Left, rs, i)
	root.Val = rs[*i]
	*i++

	fixTree(root.Right, rs, i)
}
