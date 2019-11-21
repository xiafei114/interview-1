package main

import (
	"fmt"
	"github.com/imroc/biu"
)

func main()  {
	var a, b uint8
	_ = biu.ReadBinaryString("00010111", &a)
	// 清零   按位与:两个值都为1,才为1
	// a&^a
	b = ^a
	fmt.Println(a&b)

	// 取制定位数的二进制值,    按位与, 需要的位数为1，其余位置为零, 相互与
	_ = biu.ReadBinaryString("00000111", &b)
	fmt.Println(biu.ToBinaryString(a&b))

	// 保留指定位置
	_ = biu.ReadBinaryString("00010100", &b)
	fmt.Println(biu.ToBinaryString(a&b))

	// 将某些位置为1
	_ = biu.ReadBinaryString("11100000", &b)
	fmt.Println(biu.ToBinaryString(a|b))

	date := "[00010100000100110000011000011101]"
	bs := biu.BinaryStringToBytes(date)
	fmt.Println(bs[3] >> 3 ) // 00011101
}



