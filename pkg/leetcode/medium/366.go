package medium

import "log"

/*
	366. Find Leaves of Binary Tree

	给你一棵完全二叉树，请按以下要求的顺序收集它的全部节点：

	依次从左到右，每次收集并删除所有的叶子节点
	重复如上过程直到整棵树为空
*/

func find2(n *TreeNode) [][]int {
	rss := [][]int{}
	_, rss = findleaf2(n, rss)

	return rss
}

func findleaf2(n *TreeNode, rs [][]int) (int, [][]int) {
	if n == nil {
		return 0, rs
	}

	if n.Left == nil && n.Right == nil {
		rs = append(rs, make([]int, 0))
		rs[0] = append(rs[0], n.Val)
		return 1, rs
	}

	n1, rs := findleaf2(n.Left, rs)
	n2, rs := findleaf2(n.Right, rs)

	n.Left = nil
	n.Right = nil

	dept := 0
	if n1 > n2 {
		dept = n1 + 1
	} else {
		dept = n2 + 1
	}

	if len(rs) < dept {
		rs = append(rs, make([]int, 0))
	}

	rs[dept-1] = append(rs[dept-1], n.Val)
	log.Println(rs, n.Val, dept)

	return dept, rs
}

func find1(n *TreeNode) [][]int {
	rss := [][]int{}

	for {
		rs := []int{}
		rs, ok := findLeaf(n, rs)
		rss = append(rss, rs)

		if ok {
			return rss
		}
	}
}

func findLeaf(n *TreeNode, rs []int) ([]int, bool) {
	if n == nil {
		return rs, true
	}

	if n.Left == nil && n.Right == nil {
		rs = append(rs, n.Val)
		return rs, true
	}

	r1, del1 := findLeaf(n.Left, rs)
	if del1 {
		n.Left = nil
	}

	r2, del2 := findLeaf(n.Right, rs)
	if del2 {
		n.Right = nil
	}

	rs = append(rs, r1...)
	rs = append(rs, r2...)

	return rs, false
}
