// LRU 最近最久未使用淘汰算法



package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"workPool/algorithm/links/lru/list"
)

type LRU struct {
	length uint32 
	capacity uint32 
	data map[string]interface{}
	keys []string
	mutex sync.RWMutex
}


type Node struct {
	Key string
	Value interface{}
}


func NewLRU(l uint32) *LRU {
	keys := make([]string, l, l)
	return &LRU{
		length: l,
		data: make(map[string]interface{}), 
		keys: keys,
	}
}

func(l *LRU) addOne() {
	atomic.AddUint32(&l.capacity, 1)
}

func(l *LRU) get(key string) interface{}{
	if l.keys == nil {
		return nil
	}

	l.mutex.RLock()
	defer l.mutex.RUnlock()

	v, ok := l.data[key]
	if ok {
		l.moveToFront(key)
		return v
	}

	return nil
}

func(l *LRU) Set(key string, value interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	_, ok := l.data[key]
	if ok {
		fmt.Println("hit the cache", key)
		l.data[key] = value
		l.moveToFront(key)
		return
	} 


	l.data[key] = value
	l.pushFront(key)

	if l.capacity < l.length {
		l.addOne()
	}
	return
}

func(l *LRU) removeOldKey(key string) {
	fmt.Println("remove message", key)
}

func(l *LRU) pushFront(key string) {
	for i:=l.capacity;i>0;i-- {
		// 最后一个值, 丢弃该数据
		if i == l.length {
			delete(l.data, l.keys[l.length - 1])
			continue
		}
		l.keys[i] = l.keys[i-1]
	}
	l.keys[0] = key
}

func(l *LRU) print() {
	for _, v := range l.keys{
		if v != "" {
			fmt.Printf("k: %s, v: %v;", v, l.data[v])
		}
	}

	fmt.Println()
}

func(l *LRU) moveToFront(item string) {
	var index int
	for k, v := range l.keys {
		if v == item {
			index = k 
			break
		}
	}

	if index == 0 {
		return 
	}

	for i:=index;i>0;i-- {
		l.keys[i] = l.keys[i-1]
	}
	l.keys[0] = item
} 

func main() {
	//data := []Node{
	//	{
	//		Key: "1",
	//		Value: 1,
	//	},
	//	{
	//		Key: "2",
	//		Value: 2,
	//	},
	//	{
	//		Key: "3",
	//		Value: 3,
	//	},
	//	{
	//		Key: "4",
	//		Value: 4,
	//	},
	//	{
	//		Key: "5",
	//		Value: 5,
	//	},
	//	{
	//		Key: "3",
	//		Value: 3,
	//	},
	//	{
	//		Key: "3",
	//		Value: 4,
	//	},
	//}
	//lru := NewLRU(4)
	//for _, item := range data{
	//	lru.Set(item.Key, item.Value)
	//	lru.print()
	//}
	//
	//c := []string{"3", "1", "5"}
	//
	//for _, k := range c {
	//	r := lru.get(k)
	//	fmt.Println("get cache", r)
	//}
	lru1 := list.NewLRU(3)
	_, err := lru1.Get("hello")
	if err != nil {
		fmt.Println(err)
	}
	_, _ = lru1.Set("hello", 1)
	_, _ = lru1.Set("world", 2)

	head := lru1.Head
	for {
		if head == nil{
			break
		}

		fmt.Printf("key: %s, data: %d \n", head.Key, head.Data)
		head = head.Next
	}
}
