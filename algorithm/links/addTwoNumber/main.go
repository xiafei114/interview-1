// 2.两数相加
// @links https://leetcode-cn.com/problems/add-two-numbers/
package main

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

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 5,
	}

	l2 := &ListNode{
		Val: 5,
	}

	r := addTwoNumbers(l1, l2)

	for {
		if r != nil {
			log.Println(r.Val)
			r = r.Next
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, result *ListNode
	carry, current := 0, 0
	v1, v2 := 0, 0
	var (
		n1 *ListNode
		n2 *ListNode
	)
	for {
		if l1 == nil && l2 == nil {
			if carry > 0 { // 边界检查
				head.Next = &ListNode{Val: carry}
			}
			break
		}

		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
			n1 = l1.Next
		}

		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
			n2 = l2.Next
		}

		v := v1 + v2 + carry
		current = v % 10
		carry = v / 10
		l1 = n1
		l2 = n2

		if result == nil {
			result = &ListNode{Val: current}
			head = result
		} else {
			head.Next = &ListNode{Val: current}
			head = head.Next // 指向当前节点的指针
		}
	}

	return result
}
