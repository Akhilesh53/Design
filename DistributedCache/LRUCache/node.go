package main

type Node struct {
	key  string
	val  string
	next *Node
	prev *Node
}

// NewNode creates a new Node with the given key and value.
func NewNode(key string, val string) *Node {
	return &Node{
		key:  key,
		val:  val,
		next: nil,
		prev: nil,
	}
}

// write all getter aNd setter for node
func (n *Node) GetKey() string {
	return n.key
}

func (n *Node) SetKey(key string) {
	n.key = key
}

func (n *Node) GetVal() string {
	return n.val
}

func (n *Node) SetVal(val string) {
	n.val = val
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetNext(next *Node) {
	n.next = next
}

func (n *Node) GetPrev() *Node {
	return n.prev
}

func (n *Node) SetPrev(prev *Node) {
	n.prev = prev
}
