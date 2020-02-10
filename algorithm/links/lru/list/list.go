package list

import (
	"errors"
	"sync"
)

type LRUCache struct {
	Head, tail       *LinkNode
	keyMap           map[string]*LinkNode
	lock             *sync.Mutex
	capacity, length int
}

type LinkNode struct {
	Key       string
	Data      int
	Pre, Next *LinkNode
}

var ErrNotFound = errors.New("the key not found")

func NewLRU(capacity int) *LRUCache {
	if capacity <= 0 {
		panic("capacity must be more than zero")
	}
	return &LRUCache{
		Head:     nil,
		tail:     nil,
		capacity: capacity,
		lock:     &sync.Mutex{},
		keyMap:   make(map[string]*LinkNode),
	}
}

func (l *LRUCache) Get(key string) (int, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	val := l.keyMap[key]
	if val == nil {
		return 0, ErrNotFound
	}

	l.removeToFront(val)

	return val.Data, nil
}

func (l *LRUCache) removeToFront(node *LinkNode) {
	if node == l.Head {
		return
	} else if node == l.tail {
		l.tail = l.tail.Pre
		l.tail.Next = nil
	} else {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}

	// 头插法插入
	l.Head.Pre = node
	node.Next = l.Head
	node.Pre = nil
	l.Head = node
}

func (l *LRUCache) Set(key string, val int) (bool, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.Head == nil {
		node := &LinkNode{
			Key: key,
			Data: val,
			Pre: nil,
			Next:nil,
		}
		l.Head = node
		l.tail = l.Head

		l.keyMap[key] = node
		return true, nil
	}

	node := l.keyMap[key]
	if node != nil {
		node.Data = val
		l.removeToFront(node)
	} else {
		if len(l.keyMap) >= l.capacity {
			delete(l.keyMap, l.tail.Key)
			l.tail = l.tail.Pre
			l.tail.Next = nil
		}

		tmp := &LinkNode{
			Key: key,
			Data: val,
		}

		l.keyMap[key] = tmp
		tmp.Next = l.Head
		l.Head.Pre = tmp
		l.Head = tmp
	}

	return true, nil
}