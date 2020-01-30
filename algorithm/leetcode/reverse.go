package leetcode

// 链表的逆置
// head
func Reverse(current *ListNode) *ListNode{
	var pre, next *ListNode

	for {
		if current == nil || current.Next == nil{
			current.Next = pre // 由于最后一次循环时， next == nil, 直接返回current回导致返回一个nil值
			return current
		}
		next = current.Next // next 指向下一个节点
		current.Next = pre // 当前节点指向前一个节点
		pre = current // 指向节点的指针依次向后挪动
		current = next
	}
}

func Reverse2(head *ListNode) *ListNode{
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	newHead := &ListNode{}
	for {
		if p == nil {
			break
		}

		next := p.Next
		p.Next = newHead
		newHead = p
		p = next
	}

	return newHead
}