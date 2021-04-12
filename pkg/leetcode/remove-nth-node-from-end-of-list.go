package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	pointerMap := make(map[int]*ListNode)
	i := 0
	for {
		pointerMap[i] = cur

		if cur.Next == nil {
			break
		}
		cur = cur.Next
		i++
	}
	// fmt.Printf("%+v\n", pointerMap)
	del := len(pointerMap) - n
	delNode, ok := pointerMap[del]
	if ok {
		if del-1 >= 0 {
			pointerMap[del-1].Next = delNode.Next
		} else {
			head = head.Next
		}
	}
	return head
}
