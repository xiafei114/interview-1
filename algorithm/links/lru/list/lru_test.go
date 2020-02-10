package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LRUCache(t *testing.T) {
	ast := assert.New(t)
	lru1 := NewLRU(3)
	_, err := lru1.Get("hello")
	ast.Equal(ErrNotFound, err, "")

	res, _ := lru1.Set("hello", 1)
	ast.Equal(true, res, "result not right")

	res, _ = lru1.Set("world", 2)
	ast.Equal(true, res, "result not right")

	res, _ = lru1.Set("hello", 3)

	res, _ = lru1.Set("hhhhh", 3)

	res, _ = lru1.Set("hihihi", 3)


	head := lru1.Head
	for {
		if head == nil {
			break
		}

		t.Logf("key: %s, data: %d \n", head.Key, head.Data)
		head = head.Next
	}
}