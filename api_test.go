package golinkedlist_test

import (
	"testing"

	lk "github.com/Matej-Chmel/go-linked-list"
)

func checkNextDouble(
	node *lk.DoubleLinkNode[int32], index int, val int32, t *testing.T,
) bool {
	next := node.GetNextAt(index)
	mismatch := next == nil || next.Val != val

	if mismatch {
		throw("next", index, t)
	}

	return mismatch
}

func checkNextSingle(
	node *lk.SingleLinkNode[int32], index int, val int32, t *testing.T,
) bool {
	next := node.GetNextAt(index)
	mismatch := next == nil || next.Val != val

	if mismatch {
		throw("next", index, t)
	}

	return mismatch
}

func checkPrevDouble(
	node *lk.DoubleLinkNode[int32], index int, val int32, t *testing.T,
) bool {
	prev := node.GetPrevAt(index)
	mismatch := prev == nil || prev.Val != val

	if mismatch {
		throw("prev", index, t)
	}

	return mismatch
}

func throw(direction string, index int, t *testing.T) {
	t.Errorf("Mismatch at %s index %d", direction, index)
}

func TestDoublyLinkedList(t *testing.T) {
	head := lk.CreateDoublyLinkedList[int32](1, 2, 4, 8, 16)

	if checkNextDouble(head, 0, 1, t) ||
		checkNextDouble(head, 1, 2, t) ||
		checkNextDouble(head, 2, 4, t) ||
		checkNextDouble(head, 3, 8, t) ||
		checkNextDouble(head, 4, 16, t) {
		return
	}

	last := head.GetLast()

	if last == nil {
		t.Error("Last is nil")
		return
	}

	if checkPrevDouble(last, 0, 16, t) ||
		checkPrevDouble(last, 1, 8, t) ||
		checkPrevDouble(last, 2, 4, t) ||
		checkPrevDouble(last, 3, 2, t) ||
		checkPrevDouble(last, 4, 1, t) {
		return
	}

	expected := "1 <-> 2 <-> 4 <-> 8 <-> 16"

	if str := head.String(); str != expected {
		t.Errorf("String mismatch: %s", str)
		return
	}

	if head2 := last.GetHead(); head != head2 {
		t.Error("Heads mismatch")
	}
}

func TestSinglyLinkedList(t *testing.T) {
	head := lk.CreateSinglyLinkedList[int32](1, 2, 4, 8, 16)

	if checkNextSingle(head, 0, 1, t) ||
		checkNextSingle(head, 1, 2, t) ||
		checkNextSingle(head, 2, 4, t) ||
		checkNextSingle(head, 3, 8, t) ||
		checkNextSingle(head, 4, 16, t) {
		return
	}

	last := head.GetLast()

	if last == nil {
		t.Error("Last is nil")
		return
	}

	expected := "1 -> 2 -> 4 -> 8 -> 16"

	if str := head.String(); str != expected {
		t.Errorf("String mismatch: %s", str)
	}
}
