package main

import (
	"fmt"
	"workPool/algorithm/tree"
)

func main() {
	//head := &leetcode.ListNode{
	//	Val: 1,
	//	Next: &leetcode.ListNode{
	//		Val: 2,
	//		Next: &leetcode.ListNode{
	//			Val:  3,
	//			Next: &leetcode.ListNode{
	//				Val:  4,
	//				Next: nil,
	//			},
	//		},
	//	},
	//}
	//
	//head = leetcode.Reverse(head)
	//
	//for {
	//	if head.Next == nil{
	//		return
	//	}
	//	log.Printf("current: %d \n", head.Val)
	//	time.Sleep(time.Second)
	//}

	data := []interface{}{"A", "B", "C", "D", "E", "F", "G"}
	root := tree.NewTree(data, 0, len(data))

	tree.PreOrder(root)
	fmt.Println()
	tree.InOrder(root)
	fmt.Println()
	tree.PostOrder(root)
}
