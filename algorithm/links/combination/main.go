package main

import "fmt"

func main() {
	str := "abc"
	combination(str)
}


func combination(s string) {
	data := []rune(s)
	var str []rune
	core(data, str, 0)
}
func core(s, str []rune, index int) {
	if index == len(s) {
		fmt.Println(string(str))
		return
	}

	// 选取开头的一个字符
	str = append(str, s[index])
	// 选择除掉开头的m-1个字符
	core(s, str, index + 1) //bc c
	str = str[0: len(str) - 1]
	core(s, str, index + 1)
}