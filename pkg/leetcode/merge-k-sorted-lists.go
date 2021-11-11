package leetcode

import "justest/pkg/leetcode/easy"

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	temp := lists[0]
	for i := 1; i < len(lists); i++ {
		temp = easy.MergeTwoLists(temp, lists[i])
	}

	return temp
}
