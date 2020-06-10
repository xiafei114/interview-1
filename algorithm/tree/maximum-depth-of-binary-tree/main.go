package maximum_depth_of_binary_tree

import (
	"fmt"
	"math"
	"strconv"
	"tree/tree"
)

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }

 // lt.104
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)
	return int(math.Max(float64(ld), float64(rd))) + 1
}


// lt.111
// 树的最小深度
// 根节点到叶子节点的最短路径
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 左子树为空, 右子树不为空
	if root.Left == nil && root.Right != nil {
		// 遍历右子树+当前节点
		return minDepth(root.Right) + 1
	}

	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}

	//左右子树都不为空
	l := minDepth(root.Left)
	r := minDepth(root.Right)

	// 取左右子树最小的depth
	return int(math.Min(float64(l), float64(r))) + 1
}


// lt.112
//给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
func pathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	// 叶子节点且sum减为0, 则存在这样的一条路径
	if root.Left == nil && root.Right == nil && sum - root.Val == 0 {
		return true
	}

	l := pathSum(root.Left, sum - root.Val)
	r := pathSum(root.Right, sum - root.Val)

	// r或者l有一个为true就存在
	return l || r
}


/////////////////////////////////////////////////////
//lt.437
// 双重递归
// 找出路径和等于给定数值的路径总数。
// 有点类似于一个二维数组, 需要两层遍历
// dfs表示从根节点开始遍历一遍, 获取包含该节点的所有满足条件的路径
func pathSum3(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return dfs(root, sum) + pathSum3(root.Left, sum) + pathSum3(root.Right, sum)
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	var rt int
	if sum - root.Val == 0 {
		rt = 1
	}

	return rt + dfs(root.Left, sum - root.Val) + dfs(root.Right, sum - root.Val)
}

/////////////////////////////////////////////////////
//lt.437
// pathSum前缀和解法
func PathPrefixSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	data := make(map[int]int)

	data[0] = 1
	var pathSum int
	return prefixSum(root, data, sum, pathSum)
}

func prefixSum(root *TreeNode, data map[int]int, target int, pathSum int) int {
	if root == nil {
		return 0
	}
	var res int
	// 计算当前节点的前缀和
	pathSum = pathSum + root.Val

	res = res + data[pathSum - target]
	data[pathSum] = data[pathSum] + 1

	res += prefixSum(root.Left, data, target, pathSum)
	res += prefixSum(root.Right, data, target, pathSum)

	data[pathSum] = data[pathSum] - 1

	return res
}



//lt.655
func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}

	// 计算高宽
	h := getHeight(root)
	w := (1 << h) - 1 // 2 ^ h -1

	// 构造输出结果
	ans := make([][]string, h, h)
	for i:=0;i<h;i++ {
		ans[i] = make([]string, w, w)
	}

	fill(root, ans, 0, 0, w - 1)
	fmt.Printf("%q", ans)
	return ans
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(getHeight(root.Left)), float64(getHeight(root.Right)))) + 1
}

func fill(root *TreeNode, ans [][]string, h, l, r int) {
	if root == nil {
		return
	}

	mid := (l + r) / 2
	ans[h][mid] = strconv.FormatInt(int64(root.Val), 10)
	fill(root.Left, ans, h + 1, l, mid - 1)
	fill(root.Right, ans, h + 1, mid + 1, r)
	return
}


// lt.102 二叉树层序遍历
// BFS遍历
func LevelOrder(root *tree.TreeNode) [][]interface{} {
	//if root == nil {
	//	return nil
	//}
	//return bfs(root)
	if root == nil {
		return nil
	}
	ans := make([][]interface{}, 0, 10)
	dfss(root, 0, ans)

	return ans
}

func bfs(root *tree.TreeNode) [][]interface{} {
	var current []*tree.TreeNode
	var next []*tree.TreeNode
	var ans [][]interface{}
	var h int
	current = append(current, root)
	for {
		if len(current) == 0 {
			if len(next) == 0 {
				return ans
			} else {
				current = next
				next = nil
				h++
			}
		}
		node := current[0]
		// 二维数组的行数就是树的高度
		for {
			if len(ans) <= h {
				ans = append(ans, make([]interface{}, 0))
			} else {
				break
			}
		}
		ans[h] = append(ans[h], node.Val)
		if node.Left != nil {
			next = append(next, node.Left)
		}

		if node.Right != nil {
			next = append(next, node.Right)
		}
		current = current[1:]
	}
}

func dfss (root *tree.TreeNode, depth int, ans *[][]interface{}) {
	if root == nil {
		return
	}
	// 输出节点
	if len(*ans) <= depth {
		*ans = append(*ans, make([]interface{}, 0))
	}
	(*ans)[depth] = append((*ans)[depth], root.Val)
	dfss(root.Left, depth + 1, ans)
	dfss(root.Right, depth + 1, ans)
}


// lt.110
// 判断是否是平衡二叉树
// 时间复杂度  O(n) + 2 * O(n/2) + 4 * O(n/4) + .... = O(nlogn)
func isBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	lh := getHeightV2(root.Left)
	rh := getHeightV2(root.Right)

	return math.Abs(float64(lh - rh)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeightV2(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(getHeightV2(root.Left)), float64(getHeightV2(root.Right)))) + 1
}

// 时间复杂度O(n)
func isBalancedV2(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	balance := true
	height(root, &balance)
	return balance
}

func height(root *tree.TreeNode, balance *bool) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left, balance)
	rh := height(root.Right, balance)
	if math.Abs(float64(lh - rh)) > 1 {
		*balance = false
		return -1
	}

	return int(math.Max(float64(lh), float64(rh))) + 1
}

// lt.1325
func RemoveLeafNodes(root *tree.TreeNode, target int) *tree.TreeNode {
	if root == nil {
		return nil
	}

	root.Left = RemoveLeafNodes(root.Left, target)
	root.Right = RemoveLeafNodes(root.Right, target)

	if root.Left == nil && root.Right == nil && root.Val.(int) == target {
		root = nil
		return root
	}

	return root
}