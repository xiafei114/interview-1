package main

// 自上向下理解
// 剩1个台阶 f(n-1) , 剩2个 f(n-2)

//一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
func jump(n int) int {
	if n <= 2 {
		return n
	}
	return jump(n - 1) + jump(n - 2)
}


type Node struct {
	Next *Node
	Value int
}



func ReverseRecursion(head *Node) *Node{
	// 1. 寻找终止条件 head.Next == nil
	if head == nil || head.Next == nil{
		return head
	}
	newNode := ReverseRecursion(head.Next) // 新的首节点, 此时head节点是newNode的前驱节点 A->B->head->newNode
	tmp := head.Next // newNode做为首节点保持不动, 只操作head以及, head.Next, 通过出栈回溯前一节点
	tmp.Next = head
	head.Next = nil

	return newNode
}

func main() {
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

	ReverseRecursion(node)
}