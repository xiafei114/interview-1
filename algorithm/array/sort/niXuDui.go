package sort

import (
	"fmt"
	"unsafe"
)

// [7,5,6,4]
// 5
func ReversePairs(nums []int) int {
	// 归并排序
	//[5,7]  [4,6]
	return  sortPairs(nums, 0, len(nums) - 1, 0)
}
func sortPairs(nums []int, left, right, count int) int {
	if left >= right {
		return count
	}
	mid := (left + right) / 2
	return m(nums, left, mid + 1, right, count)
}

func m(nums []int, left, mid, right, count int) int {
	leftData := make([]int, mid - left)
	rightData := make([]int, right - mid + 1)

	copy(leftData, nums[left: mid])
	copy(rightData, nums[mid: right + 1])

	var i,j int
	k := left

	for {
		if i >= len(leftData) || j >= len(rightData) {
			break
		}
		fmt.Printf("%d %d\n", leftData[i], rightData[j])
		if leftData[i] <= rightData[j] {
			nums[k] = leftData[i]
			k++
			i++
		} else {
			nums[k] = rightData[j]
			k++
			j++
			count = count + (mid + 1 + j)
		}
	}

	for {
		if i >= len(leftData) {
			break
		}
		nums[k] = leftData[i]
		k++
		i++
	}
	for {
		if j >= len(rightData) {
			break
		}

		nums[k] = rightData[j]
		count = count + j + (mid + 1)
		k++
		j++
	}
	return count
}

func MergeSort(nums []int, left, right int) {
	if left == right {
		return
	}

	// 求出中点
	mid := (left + right) / 2
	MergeSort(nums, left, mid)
	MergeSort(nums, mid + 1, right)

	merge(nums, left, mid + 1, right)
}

func merge(nums []int, left, mid, right int)  {
	leftData := make([]int, mid - left)
	rightData := make([]int, right - mid + 1)
	//填充数据, 左闭右开
	copy(leftData, nums[left: mid])
	copy(rightData, nums[mid: right + 1])

	var i, j int
	k := left

	for {
		if i >= len(leftData) || j >= len(rightData) {
			break
		}

		if leftData[i] < rightData[j] {
			nums[k] = leftData[i]
			k++
			i++
		} else {
			nums[k] = rightData[j]
			k++
			j++
		}
	}

	for {
		if i >= len(leftData) {
			break
		}
		nums[k] = leftData[i]
		k++
		i++
	}

	for {
		if j >= len(rightData) {
			break
		}

		nums[k] = rightData[j]
		k++
		j++
	}
}

func Shuffle(nums []int) {
	fmt.Printf("%p \n", nums)
	nums2 := append(nums, 1)
	fmt.Printf("%p \n", nums2)
	p := unsafe.Pointer(&nums[4])
	pp := uintptr(p)
	fmt.Println(*(*int)(unsafe.Pointer(pp)))
	ppp := uintptr(p) + 8
	fmt.Println(*(*int)(unsafe.Pointer(ppp)))

	q := unsafe.Pointer(&nums2[5])
	qq := uintptr(q)
	fmt.Println(*(*int)(unsafe.Pointer(qq)))
	fmt.Println((string)("string"))

	fmt.Println(*(*string)(unsafe.Pointer(&nums[1])))
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	//for len(nums) > 0 {
	//	n := len(nums)
	//	randIndex := r.Intn(n)
	//	nums[n-1], nums[randIndex] = nums[randIndex], nums[n-1]
	//	nums = nums[:n-1]
	//	fmt.Println(nums)
	//}
}