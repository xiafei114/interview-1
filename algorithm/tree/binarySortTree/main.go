package binarySortTree

type BinarySortTree struct {
	Root *Node
	Count int
}

type Node struct {
	Key string
	Value interface{}
	Left, Right *Node
}

func NewBst() *BinarySortTree {
	return &BinarySortTree{
		Root:  nil,
		Count: 0,
	}
}

func (b *BinarySortTree) Size() int {
	return b.Count
}

func (b *BinarySortTree) IsEmpty() bool {
	return b.Count == 0
}

func (b *BinarySortTree) Insert2(key string, value interface{}) {
	b.Root = b.insert(b.Root, key, value)
}

func (b *BinarySortTree) insert(node *Node, key string, value interface{}) *Node {
	if node == nil {
		b.Count ++
		return &Node{
			Key: key,
			Value: value,
		}
	}

	if node.Key == key {
		node.Value = value
	} else if node.Key > key {
		node.Left = b.insert(node.Left, key, value)
	} else {
		node.Right = b.insert(node.Right, key, value)
	}

	return node
}

func (b *BinarySortTree) Search(key string) *Node {
	return b.Search(key)
}

func (b *BinarySortTree) search(node *Node, key string) *Node {
	if node == nil {
		return nil
	}

	if node.Key == key {
		return node
	} else if node.Key > key {
		return b.search(node.Left, key)
	} else {
		return b.search(node.Right, key)
	}
}

func (b *BinarySortTree) Contain(key string) bool {
	return b.contain(b.Root, key)
}

func (b *BinarySortTree) contain(node *Node, key string) bool {
	if node == nil {
		return false
	}
	if node.Key == key {
		return true
	} else if node.Key < key {
		return b.contain(node.Left, key)
	} else {
		return b.contain(node.Right, key)
	}
}




//func (b *BinarySortTree) Search(data int) *common.Node {
//	p := b.Root
//	for {
//		if p != nil {
//			if p.Data == data {
//				return p
//			}
//			if p.Data > data {
//				p = p.Left
//			} else {
//				p = p.Right
//			}
//		}
//
//		return nil
//	}
//}
//
//func (b *BinarySortTree) Insert(data int) {
//	node := &common.Node{
//		Left:  nil,
//		Right: nil,
//		Data:  data,
//	}
//	p := b.Root
//	for {
//		if p == nil {
//			return
//		}
//
//		if p.Data > data {
//			if p.Left == nil {
//				p.Left = node
//				return
//			}
//			p = p.Left
//		} else {
//			if p.Right == nil {
//				p.Right = node
//				return
//			}
//			p = p.Right
//		}
//	}
//}
//
//func (b *BinarySortTree) Delete(data int) {
//	//1. 无子节点
//	p := b.Root
//	var pp *common.Node // p节点的父节点
//	for {
//		// 查找目标节点
//		if p != nil && p.Data != data {
//			pp = p
//			if data > p.Data {
//				p = p.Right
//			} else {
//				p = p.Left
//			}
//		}
//
//		if p == nil {return}
//
//		// 有两个叶子节点, 查找该节点左子树最大节点或者右子树的最小节点
//		// 该节点的最大节点一定是右节点且为叶子节点
//		if p.Right != nil && p.Left != nil {
//			var switchNode, parentNode *common.Node
//
//			parentNode = p
//			switchNode = p.Right
//			for {
//				// 当要替换的节点无右节点时为左树最大的节点
//				if switchNode.Right == nil {
//					break
//				}
//				if switchNode.Left != nil {
//					parentNode = switchNode
//					switchNode = p.Left
//				}
//			}
//
//			p.Data = switchNode.Data
//
//			// p转移至替换节点, 该节点一定是仅有一个子节点或者无子节点
//			p = switchNode
//			pp = parentNode
//		}
//
//		// 有一个叶子节点 或者无子节点
//		var child *common.Node
//		if p.Left != nil {
//			child = p.Left
//		} else if p.Right != nil {
//			child = p.Right
//		} else {
//			child = nil
//		}
//
//		// 无子节点
//		if pp == nil { // 删除的节点是根节点
//			b.Root = nil
//		} else if pp.Left == p {
//			pp.Left = child
//		} else if pp.Right == p {
//			pp.Right = p
//		}
//	}
//}