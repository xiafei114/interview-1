// 4. 寻找两个有序数组的中位数
// @links https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
// @todo 待优化
package main

import "log"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	r := findMedianSortedArrays(nums1, nums2)
	log.Println(r)
}

// 时间复杂度为O(m+n), 待优化
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	data := make([]int, 0)
	var i,j int
	for {
		// 超过长度的停止循环
		if i >= len(nums1) || j >= len(nums2) {
			break
		}

		// 从小到大排序
		if nums1[i] <= nums2[j] {
			data = append(data, nums1[i])
			i++
		} else {
			data = append(data, nums2[j])
			j++
		}
	}

	// 有一个数组没全部补充到data中
	if  i >= len(nums1) {
		data = append(data, nums2[j:]...)
	}

	if j >= len(nums2) {
		data = append(data, nums1[i:]...)
	}


	l := len(data)
	if l % 2 == 1 {
		mid := l / 2
		return float64(data[mid])
	} else {
		mid := l / 2
		return (float64(data[mid]) + float64(data[mid - 1])) / 2
	}
}
