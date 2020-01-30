// 1137. 第 N 个泰波那契数
// @links https://leetcode-cn.com/problems/n-th-tribonacci-number/
package main

func main() {
	tribonacci(5)
}

func tribonacci(n int) int {
	data := []int{0, 1, 1}
	if n <= 2 {return data[n]}
	var count int
	for i:=3;i <= n;i++ {
		count = 0
		for _, item := range data {
			count += item
		}
		data = append(data[1:3], count)
	}

	return count
}