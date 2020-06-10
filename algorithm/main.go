package main

import (
	"fmt"
	"sync"
	//maximum_depth_of_binary_tree "tree/tree/maximum-depth-of-binary-tree"
)

type LRUCache struct {
	capacity int
	keys map[int]int
	data []int
	lock *sync.Mutex
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		keys: make(map[int]int),
		data: make([]int, 0, capacity),
		lock: &sync.Mutex{},
	}
}


func (this *LRUCache) get(key int) int {
	this.lock.Lock()
	defer this.lock.Unlock()
	if v, ok := this.keys[key]; ok {
		this.moveToFront(key)
		return v
	}

	return -1
}


func (this *LRUCache) put(key int, value int)  {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.keys[key]; ok {
		this.keys[key] = value
		this.moveToFront(key)
		return
	}

	// 不存在
	if len(this.data) >= this.capacity {
		old := this.data[len(this.data) - 1]
		this.data = this.data[:len(this.data)-1]
		delete(this.keys, old)
	}

	this.keys[key] = value
	this.data = append([]int{key}, this.data...)
}


func (this *LRUCache) moveToFront(key int) {
	var value int
	for k, v:= range this.data {
		if v == key {
			value = k
			break
		}
	}

	if value != len(this.data) - 1{
		tmp := append(this.data[:value], this.data[value+1:]...)
		this.data = append([]int{key}, tmp...)
	} else {
		this.data = append([]int{key}, this.data[:len(this.data) - 1]...)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	cache := Constructor(2)
	cache.put(2, 1)
	cache.put(3,2)
	cache.get(3)       // 返回  1
	cache.get(2)       // 返回  1
	cache.put(4, 3)    // 该操作会使得密钥 2 作废
	d1 := cache.get(2)       // 返回 -1 (未找到)
	d2 := cache.get(3)       // 返回  3
	d3 := cache.get(4)       // 返回  1
	fmt.Println(d1, d2, d3)

	//head := &leetcode.ListNode{
	//	Val: 1,
	//	Next: &leetcode.ListNode{
	//		Val: 2,
	//		Next: &leetcode.ListNode{
	//			Val:  3,
	//			Next: &leetcode.ListNode{
	//				Val:  4,
	//				Next: nil,
	//			},
	//		},
	//	},
	//}
	//
	//head = leetcode.Reverse(head)
	//
	//for {
	//	if head.Next == nil{
	//		return
	//	}
	//	log.Printf("current: %d \n", head.Val)
	//	time.Sleep(time.Second)
	//}

	//data := []interface{}{"A", "B", "C", "D", "E", "F", "G"}
	//root := tree.NewTree(data, 0, len(data))
	//
	//tree.PreOrder(root)
	//fmt.Println()
	//tree.InOrder(root)
	//fmt.Println()
	//tree.PostOrder(root)

	//data := []int{1}
	//data2 := [][]int{{1}}
	//test(data, data2)
	//fmt.Println(data, data2)
	//a := []int{}
	//a = append(a, 7,8,9)
/*	a := []int{7,8,9}
	fmt.Printf("%p, %p, %p\n", &a[0], &a[1], &a[2])
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
	ap(a)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
	p := unsafe.Pointer(&a[1])
	q := uintptr(p)+8
	fmt.Println(unsafe.Pointer(q))
	t := (*int)(unsafe.Pointer(q))
	fmt.Println(*t)
 */

	//test(nil, nil)
	//data := []interface{}{1,2,3,2,nil,2,4}
	//root := tree.NewTree(data, 0, len(data))
	//maximum_depth_of_binary_tree.RemoveLeafNodes(root, 2)
}

func ap(a []int) {
	a = append(a, 10)
	fmt.Printf("%p, %p, %p, %p", &a[0], &a[1], &a[2], &a[3])
}

//func test(data []int, data2 [][]int) {
//	v := []interface{}{3,9,20,nil,nil,15,7}
//	root := tree.NewTree(v, 0, len(v))
//	vv := maximum_depth_of_binary_tree.LevelOrder(root)
//	fmt.Println(vv)
//}
