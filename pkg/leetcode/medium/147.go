package medium

/* 
	147. Insertion Sort List
*/

import (
	"log"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func init() {
	log.SetFlags(log.Lshortfile)
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 用户保存 head 的伪节点，不直接使用 head
	ptHead := &ListNode{
		Next: head,
	}
	// 已经排好序的最后一个节点
	cur := ptHead.Next

	for cur != nil {
		// 还没有排好序的第一个节点
		// 排序部分和乱序部分的分界点
		tocmp := cur.Next
		if tocmp == nil {
			break
		}

		// 遍历已排序的链表并插入 tocmp
		// 用于遍历的指针
		c := ptHead.Next
		// c的前一个节点
		ptC := ptHead

		for c != nil {
			if tocmp.Val < c.Val {
				// 插入到前面
				ptC.Next = tocmp
				tocmp.Next, cur.Next = c, tocmp.Next

				break
			}

			// 更新 c ptC，继续遍历
			ptC = c
			c = c.Next

			// 遍历结束
			if c == tocmp {
				// 更新cur
				cur = cur.Next
				break
			}
		}
	}

	return ptHead.Next
}
