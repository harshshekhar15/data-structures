package node

type Node struct {
	Val      int
	NextNode *Node
}

// New return a new instance of Node
func New(value int) *Node {
	return &Node{
		Val:      value,
		NextNode: nil,
	}
}
