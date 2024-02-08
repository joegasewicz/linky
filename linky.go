package linky

import "errors"

type List struct {
	Size      int
	Head      *Node
	Tail      *Node
	currentID int
}

// Node is the first node in the list
type Node struct {
	Next *Node
	Data *interface{}
	ID   int
}

// ListInit should be called before any operation can be used on with
// the linked list
func ListInit(data *interface{}) *List {
	n := Node{
		Data: data,
		Next: nil,
		ID:   1,
	}
	return &List{
		Size: 1,
		Head: &n,
		Tail: &n,
	}
}

// ListDestroy destroys the linked list. No operations are permitted to
// be called on the linked list after ListDestroy is called
func ListDestroy(l *List) error {
	l = nil
	if l != nil {
		return errors.New("couldn't delete List memory")
	}
	return nil
}

// InsertNext inserts a node after the current linked list element.
// If the element is nil then the new element gets placed at the head
// of the list & then return the tail.
func (l *List) InsertNext(currNode *Node, data *interface{}) (error, *Node) {
	var node *Node
	newNode := Node{
		Next: nil,
		Data: data,
		ID:   1,
	}
	if currNode.Data == l.Head.Data {
		if l.Head.Next != nil {
			err := errors.New("cannot add element to list as next element is not empty")
			return err, l.Head.Next
		}
		l.Head.Next = &newNode
		l.Size++
		l.incrementID(l.Head.Next)
		l.Tail = &newNode
		return nil, l.Head.Next
	}
	node = l.Head.Next
	for {
		if currNode.Data == node.Data || node.Next == nil {
			node.Next = &newNode
			l.Size++
			l.incrementID(&newNode)
			l.Tail = &newNode
			return nil, node.Next
		}
		node = node.Next
	}
}

// RemoveNext removes the next element after node & returns an error
// if one exists & the removed node's data. If the node is nil then the
// head of the list is removed & the head's data returned.
func (l *List) RemoveNext(node *Node) (error, *Node) {
	currNode := l.Head
	// if node is nil then remove the head & return the head's data
	if node == nil {
		headNode := l.Head
		l.Head = l.Head.Next
		l.Size--
		return nil, headNode
	}
	for {
		if currNode == node {
			if currNode.Next != nil {
				currNodeNextNode := currNode.Next
				currNode.Next = nil
				l.Size--
				return nil, currNodeNextNode
			} else {
				return errors.New("no next node"), nil
			}
		}
		currNode = currNode.Next
	}
}

// ListSize returns the total number of nodes in the list.
func (l *List) ListSize() int {
	return l.Size
}

// ListHead returns the head of the list
func (l *List) ListHead() *Node {
	return l.Head
}

// ListTail returns the tail node of the list.
func (l *List) ListTail() *Node {
	return l.Tail
}

// IsHead determines if element is the head of the list.
func (l *List) IsHead(node *Node) bool {
	if l.Head == node {
		return true
	}
	return false
}

// IsTail determines if element is the tail of the list.
func (l *List) IsTail(node *Node) bool {
	if l.Tail == node {
		return true
	}
	return false
}

// Evaluate returns the data of element's Node
func (l *List) Evaluate(node *Node) (error, *interface{}) {
	currNode := l.Head
	if currNode == node {
		return nil, currNode.Data
	}
	for {
		if currNode == node {
			return nil, currNode.Data
		}
		if currNode.Next == nil {
			break
		}
		currNode = currNode.Next
	}
	err := errors.New("cannot return data as node does not exist in list")
	return err, nil
}

// NextNode returns the next Node in the list.
func (l *List) NextNode(node *Node) (error, *Node) {
	currNode := l.Head
	err := errors.New("next node does not exist in list")
	if currNode == node {
		if currNode.Next == nil {
			return err, nil
		}
		return nil, currNode.Next
	}
	for {
		if currNode == node {
			if currNode.Next == nil {
				return err, nil
			}
			return nil, currNode.Next
		}
		if currNode.Next == nil {
			break
		}
		currNode = currNode.Next
	}
	return err, nil
}

func (l *List) incrementID(currentNode *Node) {
	l.currentID++
	currentNode.ID = l.currentID
}
