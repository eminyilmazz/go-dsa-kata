package ds

import "errors"

type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Get(idx uint) (interface{}, error) {
	if idx > l.Length()-1 {
		return nil, errors.New("index out of bounds")
	}
	current := l.head
	for i := 0; i < int(idx); i++ {
		current = current.next
	}
	return current.val, nil
}

func (l *LinkedList) Length() uint {
	if l.head == nil {
		return 0
	}
	current := l.head
	var length uint
	for current.next != nil {
		current = current.next
		length++
	}
	return length + 1
}

func (l *LinkedList) Search(needle interface{}) int {
	current := l.head
	var i int
	for current != nil {
		if current.val == needle {
			return i
		}
		i++
		current = current.next
	}
	return -1
}

func (l *LinkedList) Prepend(val interface{}) {
	node := &Node{val: val, next: nil}
	if l.head == nil {
		l.head = node
		return
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *LinkedList) Append(val interface{}) {
	node := &Node{val: val, next: nil}
	if l.head == nil {
		l.head = node
		return
	}
	current := l.head
	//traverse
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (l *LinkedList) InsertAt(idx uint, val interface{}) error {
	if idx > l.Length() {
		return errors.New("index out of bounds")
	}
	node := &Node{val: val, next: nil}

	if idx == 0 {
		l.Prepend(val)
		return nil
	}

	current := l.head
	for i := 0; i < int(idx); i++ {
		current = current.next
	}
	node.next = current.next
	current.next = node
	return nil
}

func (l *LinkedList) removeTail() error {
	if l.head == nil {
		return errors.New("empty linked list")
	}

	current := l.head
	var prev *Node
	for current.next != nil {
		prev = current
		current = current.next
	}

	// this means it never went into the for loop and those this is a single node linked list
	// thus, releasing the head
	if prev == nil {
		l.head = nil
		return nil
	}
	prev.next = nil
	return nil
}

func (l *LinkedList) removeHead() error {
	if l.head == nil {
		return errors.New("empty linked list")
	}
	l.head = l.head.next
	return nil
}

func (l *LinkedList) removeAt(idx uint) error {
	if idx >= l.Length() {
		return errors.New("index out of bounds")
	}

	if idx == 0 {
		l.head = l.head.next
		return nil
	}

	current := l.head
	var prev *Node
	for i := 0; i < int(idx); i++ {
		prev = current
		current = current.next
	}
	prev.next = current.next
	return nil
}
