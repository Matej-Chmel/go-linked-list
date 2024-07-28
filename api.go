package golinkedlist

import at "github.com/Matej-Chmel/go-any-to-string"

// A node in doubly linked list
type DoubleLinkNode[T any] struct {
	Next *DoubleLinkNode[T]
	Prev *DoubleLinkNode[T]
	Val  T
}

// Constructs a new node for doubly linked list
func NewDoubleLinkNode[T any](
	prev *DoubleLinkNode[T], next *DoubleLinkNode[T], val T,
) *DoubleLinkNode[T] {
	return &DoubleLinkNode[T]{Next: next, Prev: prev, Val: val}
}

// Constructs a new node for doubly linked list
// with nil pointers as previous and next nodes
func NewEmptyDoubleLinkNode[T any]() *DoubleLinkNode[T] {
	var val T
	return NewDoubleLinkNode(nil, nil, val)
}

// Returns true if the values of a sublist from this node to the last
// match the values of a sublist from the other node to its last node.
// To compare values an equals function is required.
func (n *DoubleLinkNode[T]) AreEqual(
	other *DoubleLinkNode[T], equals func(*T, *T) bool,
) bool {
	a, b := n, other

	if a == nil || b == nil {
		return a == nil && b == nil
	}

	for {
		if !equals(&a.Val, &b.Val) {
			return false
		}

		if a.Next == nil || b.Next == nil {
			a = a.Next
			b = b.Next
			break
		} else if a.Next.Prev != a || b.Next.Prev != b {
			return false
		}

		a = a.Next
		b = b.Next
	}

	return a == nil && b == nil
}

// Returns a string representation from this node to the last
// while respecting default format for the elements of the list
// and given format symbols for the list itself
func (n *DoubleLinkNode[T]) Format(symbols *FormatSymbols) string {
	return n.FormatCustom(at.NewOptions(), symbols)
}

// Returns a string representation from this node to the last
// while respecting given format symbols for the list itself
// and converting every node value using any-to-string library
func (n *DoubleLinkNode[T]) FormatCustom(
	options *at.Options, symbols *FormatSymbols,
) string {
	return n.FormatCustomFunction(func(t *T) string {
		return at.AnyToStringCustom(*t, options)
	}, symbols)
}

// Returns a string representation from this node to the last
// while respecting given format symbols for the list itself
// and converting every node value using a conversion function
func (n *DoubleLinkNode[T]) FormatCustomFunction(
	conv func(*T) string, symbols *FormatSymbols,
) string {
	node := &doubleImpl[T]{node: n}
	return formatToString(conv, node, symbols)
}

// Returns previous node that has nil Prev pointer
func (n *DoubleLinkNode[T]) GetHead() *DoubleLinkNode[T] {
	current := n

	for current.Prev != nil {
		current = current.Prev
	}

	return current
}

// Returns next node reached after index jumps
func (n *DoubleLinkNode[T]) GetNextAt(index int) *DoubleLinkNode[T] {
	current := n

	for i := 0; current != nil && i < index; i++ {
		current = current.Next
	}

	return current
}

// Returns last node that has nil Next pointer
func (n *DoubleLinkNode[T]) GetLast() *DoubleLinkNode[T] {
	current := n

	for current.Next != nil {
		current = current.Next
	}

	return current
}

// Returns previous node after index jumps
func (n *DoubleLinkNode[T]) GetPrevAt(index int) *DoubleLinkNode[T] {
	current := n

	for i := 0; current != nil && i < index; i++ {
		current = current.Prev
	}

	return current
}

// Returns true if for every nodes statement (node.Next.Prev == node)
// is true
func (n *DoubleLinkNode[T]) IsValid() bool {
	for n != nil && n.Next != nil {
		if n.Next.Prev != n {
			return false
		}

		n = n.Next
	}

	return true
}

// Returns a string representation from this node to the last
// in format [1 <> 2 <> 3]
func (n DoubleLinkNode[T]) String() string {
	return n.Format(NewFormatSymbols(false))
}

// Creates connected node for doubly linked list from given values
// and returns the head of that list
func CreateDoublyLinkedList[T any](values ...T) *DoubleLinkNode[T] {
	return CreateDoublyLinkedListFromSlice(values)
}

// Creates connected node for doubly linked list from given slice
// and returns the head of that list
func CreateDoublyLinkedListFromSlice[T any](values []T) *DoubleLinkNode[T] {
	dummy := NewEmptyDoubleLinkNode[T]()
	tail := dummy

	for _, val := range values {
		node := NewDoubleLinkNode(tail, nil, val)
		tail.Next = node
		tail = node
	}

	if dummy.Next != nil {
		dummy.Next.Prev = nil
	}

	return dummy.Next
}

// A node in doubly linked list
type SingleLinkNode[T any] struct {
	Next *SingleLinkNode[T]
	Val  T
}

// Constructs a new node for singly linked list
// with nil Next pointer
func NewEmptySingleLinkNode[T any]() *SingleLinkNode[T] {
	var val T
	return NewSingleLinkNode(nil, val)
}

// Constructs a new node for singly linked list
func NewSingleLinkNode[T any](next *SingleLinkNode[T], val T) *SingleLinkNode[T] {
	return &SingleLinkNode[T]{Next: next, Val: val}
}

// Returns true if the values of a sublist from this node to the last
// match the values of a sublist from the other node to its last node.
// To compare values an equals function is required.
func (n *SingleLinkNode[T]) AreEqual(
	other *SingleLinkNode[T], equals func(*T, *T) bool,
) bool {
	a, b := n, other

	for a != nil && b != nil {
		if !equals(&a.Val, &b.Val) {
			return false
		}

		a = a.Next
		b = b.Next
	}

	return a == nil && b == nil
}

// Returns a string representation from this node to the last
// while respecting default format for the elements of the list
// and given format symbols for the list itself
func (n *SingleLinkNode[T]) Format(symbols *FormatSymbols) string {
	return n.FormatCustom(at.NewOptions(), symbols)
}

// Returns a string representation from this node to the last
// while respecting given format symbols for the list itself
// and converting every node value using any-to-string library
func (n *SingleLinkNode[T]) FormatCustom(
	options *at.Options, symbols *FormatSymbols,
) string {
	return n.FormatCustomFunction(func(t *T) string {
		return at.AnyToStringCustom(*t, options)
	}, symbols)
}

// Returns a string representation from this node to the last
// while respecting given format symbols for the list itself
// and converting every node value using a conversion function
func (n *SingleLinkNode[T]) FormatCustomFunction(
	conv func(*T) string, symbols *FormatSymbols,
) string {
	node := &singleImpl[T]{node: n}
	return formatToString(conv, node, symbols)
}

// Returns last node that has nil Next pointer
func (n *SingleLinkNode[T]) GetLast() *SingleLinkNode[T] {
	current := n

	for current.Next != nil {
		current = current.Next
	}

	return current
}

// Returns next node reached after index jumps
func (n *SingleLinkNode[T]) GetNextAt(index int) *SingleLinkNode[T] {
	current := n

	for i := 0; current != nil && i < index; i++ {
		current = current.Next
	}

	return current
}

// Returns a string representation from this node to the last
// in format [1 > 2 > 3]
func (n SingleLinkNode[T]) String() string {
	return n.Format(NewFormatSymbols(true))
}

// Creates connected node for singly linked list from given values
// and returns the head of that list
func CreateSinglyLinkedList[T any](values ...T) *SingleLinkNode[T] {
	return CreateSinglyLinkedListFromSlice(values)
}

// Creates connected node for singly linked list from given slice
// and returns the head of that list
func CreateSinglyLinkedListFromSlice[T any](values []T) *SingleLinkNode[T] {
	dummy := NewEmptySingleLinkNode[T]()
	tail := dummy

	for _, val := range values {
		node := NewSingleLinkNode(nil, val)
		tail.Next = node
		tail = node
	}

	return dummy.Next
}
