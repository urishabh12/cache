package doublyLinkedList

import "errors"

type DoublyLinkedList struct {
	Head *Container
	Tail *Container
	size int
}

type Container struct {
	Value interface{}
	next  *Container
	prev  *Container
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (d *DoublyLinkedList) Append(value interface{}) error {
	node := &Container{
		Value: value,
	}

	//If linked list is empty
	if d.IsEmpty() {
		d.initialize(node)
	} else {
		node.prev = d.Tail
		d.Tail.next = node
		d.Tail = node
	}
	d.size++
	return nil
}

func (d *DoublyLinkedList) Prepend(value interface{}) error {
	node := &Container{
		Value: value,
	}

	//If linked list is empty
	if d.IsEmpty() {
		d.initialize(node)
		return nil
	}

	d.size++
	node.next = d.Head
	d.Head.prev = node
	d.Head = node

	return nil
}

func (d *DoublyLinkedList) Get(index int) (interface{}, error) {
	node := d.Head
	i := 0
	for i < index {
		node = node.next
		i++

		if node == nil {
			return nil, errors.New("index out of bound")
		}
	}

	return node.Value, nil
}

func (d *DoublyLinkedList) ListAll() []interface{} {
	node := d.Head
	resp := []interface{}{}

	for node != nil {
		resp = append(resp, node.Value)
	}

	return resp
}

func (d *DoublyLinkedList) Delete(node *Container) error {
	prev := node.prev
	next := node.next

	//When there is only one element in dll
	//When it's Head
	//When it's Tail
	//When it's in the middle
	if d.size == 1 {
		d.Head = nil
		d.Tail = nil
	} else if prev == nil {
		d.Head = next
		next.prev = nil
	} else if next == nil {
		d.Tail = prev
		prev.next = nil
	} else {
		prev.next = next
		next.prev = prev
	}

	d.size--

	return nil
}

func (d *DoublyLinkedList) IsEmpty() bool {
	return d.size == 0
}

func (d *DoublyLinkedList) initialize(node *Container) {
	d.Head = node
	d.Tail = node
}
