package golinkedlist

// Implementation of ite.CommonNode[T] for DoubleLinkNode
type doubleImpl[T any] struct {
	node *DoubleLinkNode[T]
}

func (i *doubleImpl[T]) hasNext() bool {
	return i.node.Next != nil
}

func (i *doubleImpl[T]) next() commonNode[T] {
	i.node = i.node.Next
	return i
}

func (i *doubleImpl[T]) val() T {
	return i.node.Val
}

// Implementation of ite.CommonNode[T] for SingleLinkNode
type singleImpl[T any] struct {
	node *SingleLinkNode[T]
}

func (i *singleImpl[T]) hasNext() bool {
	return i.node.Next != nil
}

func (i *singleImpl[T]) next() commonNode[T] {
	i.node = i.node.Next
	return i
}

func (i *singleImpl[T]) val() T {
	return i.node.Val
}
