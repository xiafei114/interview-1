package search

func BinarySearch(data []int, n int, target int) int {
	var left, right int = 0, n - 1
	for {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		if target == data[mid] {
			return mid
		} else if data[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

func BinarySearch2(data []int, n int, target int) int {
	return __binarySearch(data, 0, n - 1, target)
}

func __binarySearch(data []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if data[mid] == target {
		return mid
	} else if data[mid] > target {
		__binarySearch(data, left, mid - 1, target)
	} else {
		__binarySearch(data, left + 1, mid, target)
	}
}