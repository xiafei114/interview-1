package main

import (
	"fmt"
	"tree/array/dp"
)

func main()  {
	//data := []int{7,5,6,4}
	////sort.MergeSort(data, 0, len(data) - 1)
	//count := sort.ReversePairs(data)
	//fmt.Println(count)
	data := dp.MincostTickets3([]int{1,4,6,7,8,20}, []int{2,7,15})
	fmt.Println(data)
}
