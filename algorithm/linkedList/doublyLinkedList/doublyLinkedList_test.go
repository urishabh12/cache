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
