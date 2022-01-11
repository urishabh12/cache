package doublyLinkedList

import "errors"

type DoublyLinkedList struct {
	head *container
	tail *container
	size int
}

type container struct {
	value interface{}
	next  *container
	prev  *container
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (d *DoublyLinkedList) Append(value interface{}) error {
	node := &container{
		value: value,
	}

	//If linked list is empty
	if d.IsEmpty() {
		d.initialize(node)
	}

	d.size++
	d.tail.next = node
	node.prev = d.tail
	d.tail = node

	return nil
}

func (d *DoublyLinkedList) Prepend(value interface{}) error {
	node := &container{
		value: value,
	}

	//If linked list is empty
	if d.IsEmpty() {
		d.initialize(node)
		return nil
	}

	d.size++
	node.next = d.head
	d.head.prev = node
	d.head = node

	return nil
}

func (d *DoublyLinkedList) Get(index int) (interface{}, error) {
	node := d.head
	i := 0
	for i < index {
		node = node.next
		i++

		if node == nil {
			return nil, errors.New("index out of bound")
		}
	}

	return node.value, nil
}

func (d *DoublyLinkedList) ListAll() []interface{} {
	node := d.head
	resp := []interface{}{}

	for node != nil {
		resp = append(resp, node.value)
	}

	return resp
}

func (d *DoublyLinkedList) IsEmpty() bool {
	return d.size == 0
}

func (d *DoublyLinkedList) initialize(node *container) {
	d.size++
	d.head = node
	d.tail = node
}
