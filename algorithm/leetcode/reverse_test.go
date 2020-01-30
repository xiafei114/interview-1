package leetcode

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	node := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	newNode := &ListNode{}
	err := node.Copy(newNode)
	if err != nil {
		t.Errorf("copy failed, err:%v", err)
	}

	fmt.Println(newNode)
	head := Reverse(node)
	for {
			if head == nil{
				break
			}
			print(head.Val)
			head = head.Next
	}

	fmt.Println()
	head2 := Reverse2(newNode)
	for {
		if head2 == nil || head2.Next == nil{
			break
		}
		print(head2.Val)
		head2 = head2.Next
	}
}
