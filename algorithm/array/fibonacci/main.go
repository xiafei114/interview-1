package main

import (
	"fmt"
)

func main() {
	num := 9
	data := fibonacci2(num)
	data2 := fibonacci1(num)
	fmt.Println(data, data2)
}

// O(2^n) 递归算法
func fibonacci1(num int) int {
	if num <= 2 {
		return 1
	}
	// 2 + 1 + 2 + 1 = 6 
	return fibonacci1(num - 1) + fibonacci1(num - 2)
}

// O(n) 非递归算法
func fibonacci2(num int) int {
	if num <= 2 {
		return 1
	}
	n1, n2 := 1, 1
	var data int
	for i:=3;i<=num;i++ {
		data = n1 + n2
		n1 = n2
		n2 = data
	}
	
	return data
}
