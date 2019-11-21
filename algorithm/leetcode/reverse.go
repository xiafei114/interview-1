package leetcode

import "log"

// 链表的逆置
func Reverse(head *ListNode) *ListNode{
	var pre, next *ListNode

	for {
		if head.Next == nil {
			head.Next = pre
			break
		}
		next = head.Next
		head.Next = pre
		pre = head
		head = next

		log.Println(head.Val)
	}

	return head
}