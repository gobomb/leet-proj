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

func removeNthFromEndRangeTwo(head *ListNode, n int) *ListNode {
	cur := head

	l := 0
	// fmt.Printf("long: %v list: %v\n", l, head)

	// first range
	for {
		if cur == nil {
			break
		}
		if cur.Next == nil {
			l++
			break
		}
		l++

		cur = cur.Next
	}

	// get index to del
	del := l - n
	i := 0
	if del == 0 {
		return head.Next
	}
	cur = head

	// second range
	for {
		if i > l {
			return head
		}
		if del == i+1 {
			cur.Next = cur.Next.Next
			return head
		}

		if cur.Next == nil {
			break
		}
		cur = cur.Next
		i++
	}

	return head
}

func removeNthFromEndFastSlowP(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	slow := head
	fast := head
	i := 0

	for {
		if fast.Next == nil {
			break
		}
		if i == n {
			slow = slow.Next
			i = n - 1
		}
		fast = fast.Next
		i++
	}

	// 两指针的具体刚好等于n，删除慢指针的下一个node
	if i == n {
		slow.Next = slow.Next.Next
	}
	// fast已经走完，i等于节点数-1;n等于节点数，意味着需删除头节点
	// n有大于节点数的可能，这里题意假设n不大于节点数
	if i == n-1 {
		head = head.Next
	}
	return head
}
