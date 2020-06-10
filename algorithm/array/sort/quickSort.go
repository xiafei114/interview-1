package sort

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