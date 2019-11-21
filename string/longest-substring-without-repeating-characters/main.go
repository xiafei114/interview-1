// 3. 无重复字符的最长字串
// @links https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
package main

import "log"

func main() {
	s := "bbbbb"
	r := slidingWindowTwo(s)
	log.Println(r)
}



// 暴力破解法
// 遍历出字符串的所有字串, 将每个字符加入到map中, 判断对应的字符串是有已经存在过了
func violence(s string) int {
	length, max := len(s), 0

	for i:=0;i<length;i++ {
		for j:=i+1;j<=length;j++ {
			if allUnique(s, i, j) {
				max = Max(max, j - i)
			}
		}
	}

	return max
}

// 滑动窗口
func slidingWindowOne(s string) int {
	l := len(s)
	data := make(map[uint8]int)
	i, j, max := 0, 0, 0

	for {
		// 当有任意一个值索引越界后, 结束查询
		if i >= l || j >= l {
			return max
		}

		// 检查窗口中是否有制定的值
		_, ok := data[s[j]]
		if !ok {
			data[s[j]] = 0
			j++
			max = Max(max, j - i)
		} else {
			// 当发现窗口中存在重复值, 一直挪动i的位置, 直到原有i..j 中重复的索引位置过去(存在优化空间)
			delete(data, s[i])
			i++
		}
	}
}

// 滑动窗口优化版本
func slidingWindowTwo(s string) int {
	data := make(map[uint8]int)
	var i,j, max int
	l := len(s)

	for {
		if i >= l || j >= l {
			return max
		}

		index, ok := data[s[j]]
		if !ok {
			// 存储索引位置
			data[s[j]] = j
			j++
			max = Max(max, j - i)
		} else {
			// 如果当前的index比i要小, 说明数据落在了窗口之外
			if index < i {
				// 原先的数据进行修改
				data[s[j]] = j
				j++
				max = Max(max, j - i)
			} else {
				// i直接索引到index+1的位置, 减少了 i..j` 的重复判断
				i = index + 1
			}
		}
	}
}

func allUnique(s string, start, end int) bool {
	data := map[uint8]int{}
	for i:=start;i<end;i++ {
		char := s[i]
		if _ ,ok := data[char]; ok {
			return false
		}
		data[char] = 0
	}

	return true
}
func Max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
