package dp

// dp[i][j] 表示矩阵中当前元素组成最大正方形的周长
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	if len(matrix[0]) > 0 {
		for i:=0;i<len(matrix[0]);i++ {
			if matrix[0][i] == 1 {
				return 1
			}
		}
	}

	dp := make([][]int, len(matrix))
	var ans int
	for i:=0;i<len(matrix);i++ {
		dp[i] = make([]int, len(matrix[i]) + 1)
		for j:=0;j<len(matrix[i]);j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[0][j] == 1 || dp[i][0] == 1 {
				ans = 1
			}
		}
	}

	for i:=1;i<len(matrix);i++ {
		for j:=1;j<len(matrix[i]);j++ {
			// md 表示矩阵中当前元素组成最大正方形的周长
			md := min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			// 此时的dp[i][j] 表示的是当前matrix数组的值
			if dp[i][j] > 0 && md > 1 {
				dp[i][j] = md
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans * ans
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

// dp[i][j] 表示最大边长
func maximalSquare2(matrix [][]byte) int {
	// dp的长度定义成len(matrix)+1是想向外扩一圈, 避免特殊处理
	dp := make([][]int, len(matrix) + 1)
	dp[0] = make([]int, len(matrix) + 1)
	maxSize := 0
	// 将matrix转成int型
	for i:=0;i<len(matrix) + 1;i++ {
		dp[i+1] = make([]int, len(matrix[i]) + 1)
		for j:=0;j<len(matrix[i]);j++{
			if int(matrix[i][j] - '0') == 1 {
				dp[i+1][j+1] =  min(min(dp[i+1][j], dp[i][j+1]), dp[i+1][j+1]) + 1
				if dp[i+1][j+1] > maxSize {
					maxSize = dp[i+1][j+1]
				}
			}
		}
	}

	return maxSize * maxSize
}
