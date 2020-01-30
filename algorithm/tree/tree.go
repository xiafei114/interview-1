package tree

import "fmt"

// 完全二叉树
// 假设父节点的编号为n, 若左孩子节点存在, 则左孩子节点编号为2n, 后孩子节点编号为2n+1
// 数据索引为0~n-1, 树为1～n,

// 通过数据构建二叉树
type TreeNode struct {
	Val interface{}
	Left *TreeNode
	Right *TreeNode
}

type TreeString []string

func (t *TreeString) Reset() {
	t = nil
}

var root *TreeNode

func NewTree(data []interface{}, index, length int) *TreeNode {
	tmp := make([]interface{}, 1, len(data) + 1)
	if index == 0 {
		tmp[0] = nil
		data = append(tmp, data...)
		index = 1
		length = length + 1
	}
	if index >= length || data[index] == nil{
		return nil
	}

	root := &TreeNode{Val: data[index]}
	root.Left = NewTree(data, 2 * index, length)
	root.Right = NewTree(data, 2 * index + 1, length)
	return root
}

// 前序
func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%v", root.Val)
	if root.Left != nil {PreOrder(root.Left)}
	if root.Right != nil {PreOrder(root.Right)}
}

// 中序
func InOrder(root *TreeNode) {
	if root == nil {return}
	if root.Left != nil {
		InOrder(root.Left)
	}
	fmt.Printf("%v", root.Val)
	if root.Right != nil {
		InOrder(root.Right)
	}
}

// 后序
func PostOrder(root *TreeNode) {
	if root == nil {return}
	if root.Left != nil {PostOrder(root.Left)}
	if root.Right != nil {PostOrder(root.Right)}

	fmt.Printf("%v", root.Val)
}

