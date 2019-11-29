package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 5, 6}
	limit := 10
	r := max(data, limit)
	fmt.Println(r)
}



func max(data []int, limit int) int {
	length := len(data)
	i, j := 0, length - 1
	max := 0 
	for {
		if i >= j {
			return max 
		}

		sum := data[i] + data[j]
		if sum == limit {
			return sum
		}

		if sum < limit {
			i++
		}

		if sum > limit {
			j--
			continue
		}

		if sum > max {
			max = sum
		}
	}

	return max
}

func quickSort(data []int, s, e int) {
	i, j := s, e
	x := data[i] // pivot

	for {
		if i >= j {
			break
		}

		for {
			if i < j && data[j] >= x {
				j--
			} else {
				data[i] = data[j]
				i++
				break
			}
		}

		for {
			if i < j && data[i] < x {
				i++
			} else {
				data[j] = data[i]
				j--
				break
			}
		}
	}

	data[i] = x
	quickSort(data, s, i - 1)
	quickSort(data, i + 1, e)
}