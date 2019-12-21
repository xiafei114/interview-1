// 两条链表是否有交叉点
// 链表当存在公共节点时,  从该节点开始,后续节点其实为同一条链

package main

import "fmt"

type LinkNode struct {
	Val int
	Next *LinkNode
}

func main() {
	l1 := &LinkNode{
		Val: 1,
		Next: &LinkNode{
			Val: 2,
			Next: &LinkNode{
				Val: 3,
				Next: &LinkNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	l2 := &LinkNode{
		Val: 2,
		Next: &LinkNode{
			Val: 6,
			Next: &LinkNode{
				Val: 3,
				Next: nil,
			},
		},
	}

	r := cross(l1, l2)
	fmt.Println(r)
}

// @todo 两条链表连接, 检查是否存在有环
func cross2(l1, l2 *LinkNode) bool {
	return true
}


// 将长的链条指针挪到跟短链条同样长度的位置, 在对两条链表进行数值的比较
func cross(l1, l2 *LinkNode) bool {
	if l1 == nil || l2 == nil {
		return false
	}
	len1 := l1.Len()
	len2 := l2.Len()

	var long, short *LinkNode
	var interval int

	if len1 > len2 {
		long = l1
		short = l2
		interval = len1 - len2
	} else {
		long = l2
		short = l1
		interval = len2 - len1
	}

	// 将长的链条指针挪到跟短链条同样长度的位置
	for i:=0;i<interval;i++ {
		long = long.Next
	}

	for {
		// 如果某一个链表为空了还没有找到，则不存在分叉
		if long == nil || short == nil {
			return false
		}
		// 查找到交叉点
		if long.Val == short.Val {
			return true
		}

		// 指针向后挪
		long = long.Next
		short = short.Next
	}




}


// 计算链表的长度
func (l *LinkNode) Len() int {
	p := l
	var count int
	for {
		if p == nil {
			return count
		}

		count ++
		p = p.Next
	}
}
