package main

import (
	"fmt"
	"math"
)

// @links https://leetcode-cn.com/problems/unique-paths/
// 1. 假设每个点的路径为d[i][j]
// 2. 计算位于(m,n) 有多少种走法
func uniquePaths(m int, n int) int {
	dp := make([][]int, m, m)
	for index, _ := range dp {
		dp[index] = make([]int, n, n)
	}
	// i 或者 j 为零时，只能沿着墙一种走法
	for i:=0;i<m;i++ {
		dp[i][0] = 1
	}

	for j:=0;j<n;j++ {
		dp[0][j] = 1
	}

	// 到(i,j)坐标, 有d[i-1][j], i[i][j-1]两种走法
	for i:=1;i<m;i++ {
		for j:=1;j<n;j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 最短路径和
// 设d[i][j] 是到(i,j) 节点的最小路径
// 则d[m-1][n-1] 就是求解的最小路径
// 到这个节点有两种走法(i, j-1), (i-1, j),从两种走法种选择最小的, 但是需要之前的状态
// dp[i][j] = min(dp[i][j-1] + dp[i-1][j]) + arr[i][j]
func minPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}
	m := len(grid) // 行
	n := len(grid[0]) // 列

	dp := make([][]int, m, m) // 1. 定义数组元素的含义
	for index, _ := range dp {
		dp[index] = make([]int, n, n)
	}
	// m = 2 n = 3
	// 2. 初始值
	dp[0][0] = grid[0][0]
	for i:=1;i<m;i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j:=1;j<n;j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i:=1;i<m;i++ {
		for j:=1;j<n;j++ { // 3.数组元素间的表达式
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}

	fmt.Println(dp)
	return dp[m-1][n-1]
}

// 走过的路径直接覆盖原来grid的值, 时间复杂度O(1)
func minPathSum2(grid [][]int) int {
	if grid == nil {
		return 0
	}
	m := len(grid) // 行
	n := len(grid[0]) // 列
	// m = 2 n = 3
	// 2. 初始值
	for i:=1;i<m;i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for j:=1;j<n;j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}

	for i:=1;i<m;i++ {
		for j:=1;j<n;j++ { // 3.数组元素间的表达式
			grid[i][j] = int(math.Min(float64(grid[i-1][j]), float64(grid[i][j-1]))) + grid[i][j]
		}
	}

	return grid[m-1][n-1]
}


// 买卖股票的最佳时机
// @https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/submissions/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, profit := prices[0], 0
	for i:=1;i<len(prices);i++ {
		// 计算买入股票时的价格
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] - minPrice > profit { //计算最大的利润值
			profit = prices[i] - minPrice
		}
	}

	return profit
}


// 爬楼底动态规划做法
// @Links https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	// d[i] = d[i-1] + d[i-2] 到第i个台阶时的有d[i]个走法
	// d[1] = 1 d[2] = 2
	if n <= 2 {
		return n
	}

	dp := make([]int, n, n)
	dp[0] = 1
	dp[2] = 2
	for i:=3;i<n;i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

// 爬楼梯递归做法
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}

	data := []int{1, 2}
	i := 3
	var sum int
	for {
		if i > n {
			break
		}

		sum = data[0] + data[1]
		data = []int{data[1], sum}

		i++
	}

	return sum
}


// dp[i] 在某间屋子能够达到的最大金额
// dp[i] = dp[i-2] + arr[i]
// dp[0] = arr[0]
// dp[1] = arr[1]
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}

	dp := make([]int, length, length)
	dp[0], dp[1] = nums[0], int(math.Max(float64(nums[0]), float64(nums[1])))
	// 2, 2, 1, 2
	for i:=2;i<length;i++ {
		// 两种case
		// 1. 当前轮 + 前两轮的值
		// 2. 当前轮不去, 则当前dp的上一次的dp值
		dp[i] = int(math.Max(float64(dp[i-2]) + float64(nums[i]), float64(dp[i-1])))
	}
	return dp[len(dp) - 1]
}