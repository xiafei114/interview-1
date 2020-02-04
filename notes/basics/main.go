package main

import "fmt"

func main() {
	test1()
}

// 原码, 补码, 反码
func test1() {
	var a int8 = -128
	var b = a / -1 // 128-> 补码表示 0 1000 0000 舍弃掉最高位 = 1000 0000 = -128
	fmt.Println(b)

	//var c int8 = 128
	//fmt.Println(c)
}
