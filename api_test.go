package golinkedlist_test

import (
	"fmt"
	"runtime"
	"testing"

	at "github.com/Matej-Chmel/go-any-to-string"
	lk "github.com/Matej-Chmel/go-linked-list"
)

// Wrapper around the test state
type tester struct {
	failed bool
	*testing.T
}

func newTester(t *testing.T) *tester {
	return &tester{failed: false, T: t}
}

func (t *tester) throw(skip int, format string, data ...any) {
	if t.failed {
		return
	}

	_, _, line, ok := runtime.Caller(skip)

	if ok {
		format = fmt.Sprintf("(line %d) %s", line, format)
	}

	t.Errorf(format, data...)
	t.failed = true
}

func (t *tester) throwMismatch(actual, expected string) {
	if t.failed {
		return
	}

	t.throw(2, "\n\n%s\n\n!=\n\n%s", actual, expected)
}

func checkNextDouble[T comparable](
	node *lk.DoubleLinkNode[T], index int, expected T, t *tester,
) {
	if t.failed {
		return
	}

	next := node.GetNextAt(index)

	if next == nil {
		var actual T
		throwIndex(t, "next", index, actual, expected)
	} else if next.Val != expected {
		throwIndex(t, "next", index, next.Val, expected)
	}
}

func checkNextSingle[T comparable](
	node *lk.SingleLinkNode[T], index int, expected T, t *tester,
) {
	if t.failed {
		return
	}

	next := node.GetNextAt(index)

	if next == nil {
		var actual T
		throwIndex(t, "next", index, actual, expected)
	} else if next.Val != expected {
		throwIndex(t, "next", index, next.Val, expected)
	}
}

func checkPrevDouble[T comparable](
	node *lk.DoubleLinkNode[T], index int, expected T, t *tester,
) {
	if t.failed {
		return
	}

	prev := node.GetPrevAt(index)

	if prev == nil {
		var actual T
		throwIndex(t, "prev", index, actual, expected)
	} else if prev.Val != expected {
		throwIndex(t, "prev", index, prev.Val, expected)
	}
}

func throwIndex[T any](t *tester, direction string, index int, actual, expected T) {
	t.throw(3, "Mismatch at %s index %d, %v != %v",
		direction, index, actual, expected)
}

func TestDoublyLinkedList(ot *testing.T) {
	t := newTester(ot)
	head := lk.CreateDoublyLinkedList[float64](1.1, 2.01, 4.31, 8.901, 16.32)

	checkNextDouble(head, 0, 1.1, t)
	checkNextDouble(head, 1, 2.01, t)
	checkNextDouble(head, 2, 4.31, t)
	checkNextDouble(head, 3, 8.901, t)
	checkNextDouble(head, 4, 16.32, t)

	last := head.GetLast()

	if last == nil {
		t.throw(1, "Last is nil")
	}

	checkPrevDouble(last, 0, 16.32, t)
	checkPrevDouble(last, 1, 8.901, t)
	checkPrevDouble(last, 2, 4.31, t)
	checkPrevDouble(last, 3, 2.01, t)
	checkPrevDouble(last, 4, 1.1, t)

	expected := "[1.1 <> 2.01 <> 4.31 <> 8.901 <> 16.32]"

	if actual := head.String(); actual != expected {
		t.throwMismatch(actual, expected)
	}

	if head2 := last.GetHead(); head != head2 {
		t.throw(1, "Heads mismatch")
	}

	expected = "(1.1, 2.0, 4.3, 8.9, 16.3)"
	symbols := lk.NewFormatSymbols(true)
	symbols.Start = "("
	symbols.Sep = ", "
	symbols.End = ")"

	options := at.NewOptions()
	options.FloatDecimalPlaces = 1

	if actual := head.FormatCustom(options, symbols); actual != expected {
		t.throwMismatch(actual, expected)
	}
}

func TestEmptyList(ot *testing.T) {
	t := newTester(ot)
	n1 := lk.CreateSinglyLinkedList[*lk.SingleLinkNode[bool]](nil)
	n2 := lk.CreateDoublyLinkedList[*lk.DoubleLinkNode[complex128]](nil)
	expected := "[nil]"

	if actual := n1.String(); actual != expected {
		t.throwMismatch(actual, expected)
	}

	if actual := n2.String(); actual != expected {
		t.throwMismatch(actual, expected)
	}
}

func TestSinglyLinkedList(ot *testing.T) {
	t := newTester(ot)
	head := lk.CreateSinglyLinkedList[int32](1, 2, 4, 8, 16)

	checkNextSingle(head, 0, 1, t)
	checkNextSingle(head, 1, 2, t)
	checkNextSingle(head, 2, 4, t)
	checkNextSingle(head, 3, 8, t)
	checkNextSingle(head, 4, 16, t)

	last := head.GetLast()

	if last == nil {
		t.throw(1, "Last is nil")
	}

	expected := "[1 > 2 > 4 > 8 > 16]"

	if actual := head.String(); actual != expected {
		t.throwMismatch(actual, expected)
	}

	symbols := lk.NewFormatSymbols(true)
	symbols.Start = "$$ "
	symbols.Sep = " - "
	symbols.End = " $$"

	expected = "$$ 1 - 2 - 4 - 8 - 16 $$"

	if actual := head.Format(symbols); actual != expected {
		t.throwMismatch(actual, expected)
	}
}
