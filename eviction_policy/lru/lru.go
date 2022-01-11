package lru

import (
	d "github.com/urishabh12/cache/algorithm/linkedList/doublyLinkedList"
)

type LRU struct {
	dll       *d.DoublyLinkedList
	keyToNode map[interface{}]*d.Container
}

func NewLRUEviction() *LRU {
	return &LRU{
		dll:       d.NewDoublyLinkedList(),
		keyToNode: make(map[interface{}]*d.Container),
	}
}

func (l LRU) KeyUsed(key interface{}) error {
	if l.keyToNode[key] != nil {
		l.dll.Delete(l.keyToNode[key])
	}

	l.dll.Append(key)
	l.keyToNode[key] = l.dll.Tail

	return nil
}

func (l LRU) EvictKey() (interface{}, error) {
	nodeToDelete := l.dll.Head
	l.dll.Delete(nodeToDelete)
	delete(l.keyToNode, nodeToDelete.Value)

	return nodeToDelete.Value, nil
}
