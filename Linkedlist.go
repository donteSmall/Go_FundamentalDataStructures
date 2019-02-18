package main

// "Linked List implementation"
type List struct {
	head *Node
	tail *Node
}

// Returns the head of the list !
func (l *List) First() *Node {
	return l.head
}

// Pushes a node onto the list
func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		//When assiging new element to list,link it back to node
		l.tail.next = node
		//Doubly linked list componet
		node.prev = l.tail
	}
	l.tail = node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}
type error interface {
	Error() string
}

// Returns the next node
func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}
func (l *List) Last() *Node {
	return l.tail

}
func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	_head := l.First()
	for {
		println(_head.value)
		_head = _head.Next()
		if _head == nil {
			break
		}
	}

}
