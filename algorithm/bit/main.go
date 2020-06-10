package main

import "fmt"

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

	m, n  := isEven(1),	isEven(2)
	fmt.Println(m, n)

	fmt.Println(reverse(1), reverse(-1))
	count := countOne(5)
	fmt.Println(count)
}

func swap(a, b *int) {
	*a ^= *b // a = a^b
	*b ^= *a // b = b^b^a =a
	*a ^= *b // a = a^b = a^b^a = b
}

func isEven(a int) bool {
	return 0 == (a&1)
}



func reverse(a int) int {
	// 复数补码表示
	// 000001
	// 111110 + 1 = 111111 100000 + 1 = -1

	// 000010
	// 111101 + 1 = 111110 = 100001 + 1 = 100010 = -2

	// 111111 -> 00000 + 1 = 000001 = 1
	// 100001 -> 补码处理
	// 011110 + 1 = 011111 =
	return  ^a + 1
}

func abs(a int) int {
	// 整数右移31位=0, 负数右移31位=oxffffffff -> 补码表示 -1
	i := a >> 31
	if i == 0 {
		return a
	}
	return reverse(a)
}

func abs2(a int) int {
	i := a >> 31
	return (a^i) - i
}

func highLowSwap(a uint16) uint16 {
	return a >> 8 | a << 8
}

func countOne(a int) (count int) {
	for a > 0 {
		a = a & (a - 1)
		count++
	}

	return
}