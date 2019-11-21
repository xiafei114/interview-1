package main

import (
	"workPool/algorithm/leetcode"
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

	leetcode.CalcDays()
}
