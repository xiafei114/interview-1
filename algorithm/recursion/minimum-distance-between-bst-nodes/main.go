package main

import (
	"math"
)

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

var min int

func minDiffInBST(root *TreeNode) int {
	var left, right int
	var lg, rg bool // 标记节点是否有值
	if root.Left != nil {
		data := minDiffInBST(root.Left)
		left = int(math.Abs(float64(root.Val - data)))
		lg = true
	}

	if root.Right != nil {
		data := minDiffInBST(root.Right)
		right = int(math.Abs(float64(root.Val - data)))
		rg = true
	}

	if lg == false && rg == false {
		return root.Val
	}

	if lg == true && rg == false {
		if left < min {min=left}
		return left
	} else if lg == false && rg == true {
		if right < min {min=right}
		return right
	} else {
		if left > right {
			return right
		}
		return left
	}
}