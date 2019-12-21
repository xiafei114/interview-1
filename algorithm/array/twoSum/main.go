//  leetcode 2.两数之和
// @links https://leetcode-cn.com/problems/two-sum/


package main

import "fmt"

func main() {
	nums := []int{3,2,4}
	target := 6

	r := twoSum(nums, target)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	data := make(map[int]int)

	for k, v := range nums {
		ano := target - v
		index, ok := data[ano]
		if ok {
			return []int{k, index}
		}

		data[v] = k
	}

	return []int{}
}