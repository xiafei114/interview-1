package main

import "testing"

func Test_Reverse_Recursion(t *testing.T) {
	node := &Node{
		Next:  &Node{
			Next:  &Node{
				Next:  &Node{
					Next:  nil,
					Value: 4,
				},
				Value: 3,
			},
			Value: 2,
		},
		Value: 1,
	}

	node = ReverseRecursion(node)
	for {
		if node == nil {
			break
		}

		t.Log(node.Value)
		node = node.Next
	}
}