package common_tree

type Node struct {
	Left, Right *Node
	Data int
}

func (n *Node) Set(data int) {
	n.Data = data
}