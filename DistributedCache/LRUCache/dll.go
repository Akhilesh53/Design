package main

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
	}
}

// write all getter aand setter for dll
func (dll *DoublyLinkedList) GetHead() *Node {
	return dll.head
}

func (dll *DoublyLinkedList) SetHead(head *Node) *Node {
	dll.head = head
	return dll.head
}

func (dll *DoublyLinkedList) GetTail() *Node {
	return dll.tail
}

func (dll *DoublyLinkedList) SetTail(tail *Node) *Node {
	dll.tail = tail
	return dll.tail
}

// add anode to dll
func (dll *DoublyLinkedList) Add(node *Node) {
	// check if head is nil
	if dll.GetHead() == nil {
		dll.SetHead(node)
		dll.SetTail(node)
		return
	}
	// add node to the end of the list
	dll.GetTail().SetNext(node)
	node.SetPrev(dll.GetTail())

	// move tail ptr to last
	dll.SetTail(node)
}

// MoveToLast moves the given node to the last position in the doubly linked list.
// If the node is already at the end, it returns the node itself.
// If the node is at the start, it shifts the head to the next node and adds the node to the end of the list.
// If the node is in between the head and tail, it removes the node from its current position and adds it to the end of the list.
// It returns the new tail of the list.
func (dll *DoublyLinkedList) MoveToLast(node *Node) *Node {

	// if node is already at the end
	if dll.GetTail() == node {
		return node
	}

	// if node is at the start
	if dll.GetHead() == node {
		temp := node
		//shift head to next node
		dll.SetHead(dll.GetHead().GetNext())
		dll.GetHead().SetPrev(nil)

		// add node to the end of the list
		dll.GetTail().SetNext(temp)
		temp.SetPrev(dll.GetTail())
		dll.SetTail(temp)
		dll.GetTail().SetNext(nil)
		return dll.GetTail()
	}

	// if node is in b/w of head and tail
	node.GetPrev().SetNext(node.GetNext())
	node.GetNext().SetPrev(node.GetPrev())
	dll.GetTail().SetNext(node)
	node.SetPrev(dll.GetTail())
	dll.SetTail(node)
	dll.GetTail().SetNext(nil)

	return dll.GetTail()
}

func (dll *DoublyLinkedList) PrintLL() {
	node := dll.GetHead()
	for node != nil {
		println("key :: ", node.GetKey(), "val :: ", node.GetVal())
		node = node.GetNext()
	}
}
