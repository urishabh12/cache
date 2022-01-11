package doublyLinkedList

import (
	"fmt"
	"testing"
)

func Test_CheckIfSequenceCorrectAppend(t *testing.T) {
	dll := NewDoublyLinkedList()
	values := []int{1, 2, 3, 4}
	for i := 0; i < len(values); i++ {
		dll.Append(values[i])
	}

	for i := 0; i < len(values); i++ {
		val, err := dll.Get(i)
		handleErr(t, err)
		assertEqualInt(t, values[i], val.(int))
	}
}

func Test_DeleteNodeMiddle(t *testing.T) {
	dll := NewDoublyLinkedList()
	values := []int{1, 2, 3, 4}
	for i := 0; i < len(values); i++ {
		dll.Append(values[i])
	}

	dll.Delete(dll.Head.next)
	val, err := dll.Get(1)
	handleErr(t, err)
	assertEqualInt(t, values[2], val.(int))
}

func Test_DeleteNodeHead(t *testing.T) {
	dll := NewDoublyLinkedList()
	values := []int{1, 2, 3, 4}
	for i := 0; i < len(values); i++ {
		dll.Append(values[i])
	}

	dll.Delete(dll.Head)
	val, err := dll.Get(0)
	handleErr(t, err)
	assertEqualInt(t, values[1], val.(int))
}

func Test_DeleteNodeTail(t *testing.T) {
	dll := NewDoublyLinkedList()
	values := []int{1, 2, 3, 4}
	for i := 0; i < len(values); i++ {
		dll.Append(values[i])
	}

	dll.Delete(dll.Tail)
	_, err := dll.Get(3)
	if err == nil {
		t.Fatalf("Tail not deleted")
	}
	val, err := dll.Get(2)
	handleErr(t, err)
	assertEqualInt(t, values[2], val.(int))
}

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertEqualInt(t *testing.T, a int, b int) {
	if a != b {
		t.Fatalf(fmt.Sprintf("Values not equal %d != %d", a, b))
	}
}
