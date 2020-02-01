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

// 1. len(X) = [len(O), len(O)+1]
// 2. len(X) + len(O) <= 9
// 3. X获胜 len(X) = len(O) + 1, O获胜len(X) = len(O)


func validTicTacToe(board []string) bool {
	var xCount,oCount,blankCount int
	for _, line := range board {
		for _, ch := range line {
			switch string(ch) {
			case "O":
				oCount ++
			case "X":
				xCount ++
			case " ":
				blankCount ++
			default:
				return false
			}
		}
	}

	if xCount + oCount + blankCount > 9 {
		return false
	}

	if xCount != oCount && xCount != oCount + 1 {
		return false
	}

	win := func(ch uint8) bool {
		for i:=0;i<3;i++ {
			if board[0][i] == ch && board[1][i] == ch && board[2][i] == ch {return true}
			if board[i][0] == ch && board[i][1] == ch && board[i][2] == ch {return true}
		}

		if ch == board[0][0] && ch == board[1][1] && ch == board[2][2] {return true}
		if ch == board[0][2] && ch == board[1][1] && ch == board[2][0] {return true}
		return false
	}

	x := "X"
	o := "O"
	if win(uint8(x[0])) && xCount != oCount + 1 {return false}
	if win(uint8(o[0])) && oCount != xCount {return false}

	return true
}

