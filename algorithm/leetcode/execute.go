package leetcode

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
	"unicode"
)

// 判断是否是回文串
func IsPalindrome(s string) bool {
	var str string
	for _, value := range s {
		if unicode.IsDigit(value) || unicode.IsLetter(value) {
			str = str + strings.ToLower(string(value))
		}
	}
	var l, r, mid int

	log.Println("字符串总长度:", len(str))
	if len(str) % 2 == 1 {
		mid = len(str) / 2
		l, r = mid - 1 , mid + 1
	} else {
		mid = (len(str) / 2) - 1
		l, r = mid, mid + 1
	}

	for {
		if l < 0 || r > len(str) {
			return true
		}

		log.Printf("[%d], [%d]\n", str[l], str[r])
		if str[l] == str[r] {
			l --
			r ++
		} else {
			return false
		}
	}
	log.Println("success")
	return true
}

// 双指针
func IsPalindromeV2(s string) bool {
	//
	isalnum := func(c rune) bool {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			return true
		}
		return false
	}
	l, r := 0, len(s) - 1
	for {
		if l >= r {
			return true
		}

		c1 := rune(s[l])
		c2 := rune(s[r])
		if isalnum(rune(s[l])) && isalnum(rune(s[r])) { // 检查两个是否都是数字或者字符
			if strings.ToLower(string(c1)) == strings.ToLower(string(c2)) {
				l++
				r--
			} else {
				return false
			}
		} else {
			if !isalnum(c1) {
				l++
			}

			if !isalnum(c2) {
				r--
			}
		}
	}
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type DeepCopy interface {
	Copy(dst interface{}) error
}

// 深拷贝
func (l *ListNode) Copy(dst interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(*l)
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}


type ListNode struct {
	Val int
	Next *ListNode
}

// 链表数据插入数组比较
func IsPalindromeV3(head *ListNode) bool {
	p := head
	var data []int
	for {
		if p == nil {
			break
		}

		data = append(data, p.Val)

		p = p.Next
	}

	for i, j := 0, len(data) - 1; i <= j; {
		if data[i] != data[j] {
			return false
		}

		i++
		j--
	}

	return true
}

