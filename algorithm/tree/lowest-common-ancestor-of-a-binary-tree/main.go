package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
				Left:  &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{
					Val:   10,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   11,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	
	node := action2(tree, &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   11,
		Left:  nil,
		Right: nil,
	})

	if node != nil {
		fmt.Println(node.Val)
	}
}


var ans *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	recurseTree(root, p, q)
	return ans
}

func recurseTree (currentNode, p, q *TreeNode) bool {
	if currentNode == nil {
		return false
	}

	var left, right, mid int

	if recurseTree(currentNode.Left, p, q) {
		left = 1
	}

	if recurseTree(currentNode.Right, p, q) {
		right = 1
	}

	if currentNode.Val == p.Val || currentNode.Val == q.Val {
		mid = 1
	}

	if mid + left + right >= 2 {
		ans = currentNode
	}

	return (mid + left + right) > 0
}


func action2(root, p, q *TreeNode) *TreeNode {
	// 将节点入栈
	stack := make([]*TreeNode, 1)
	// 存放各节点的父亲节点之间的关系
	parent := make(map[int]*TreeNode)


	parent[root.Val] = nil
	stack = append(stack, root)

	for {
		if !contain(parent, p) || !contain(parent, q) {

			node := stack[:len(stack) - 1][0]

			if node.Left != nil {
				stack = append(stack, node.Left)
				parent[node.Left.Val] = node
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
				parent[node.Right.Val] = node
			}
			// 查左边的节点
		} else {
			break
		}
	}

	pointer := p
	var ancestors map[int]*TreeNode
	for {
		if pointer != nil {
			ancestors[pointer.Val] = pointer
			pointer = parent[pointer.Val] // 向上遍历父结点
		} else {
			break
		}
	}

	// p节点的祖先节点数据构建完成
	// 查看q节点的祖先节点有没有跟p节点重合的

	for {
		if q == nil {
			return nil
		}

		if !contain(ancestors, q) {
			q = parent[q.Val]
		} else {
			return q
		}
	}
}

func contain(parent map[int]*TreeNode, node *TreeNode) (ok bool) {
	if _, ok := parent[node.Val]; ok {
		return true
	}

	return
}