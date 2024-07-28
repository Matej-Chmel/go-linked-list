package golinkedlist

import (
	"strings"

	at "github.com/Matej-Chmel/go-any-to-string"
)

// Interface that represents one singly or doubly linked node
type commonNode[T any] interface {
	// Returns true if next node is available
	hasNext() bool
	// Returns the next node
	next() commonNode[T]
	// Returns the value of the node
	val() T
}

// Format symbols for converting a linked list to a string
type FormatSymbols struct {
	// Symbol at the end of the list
	End string
	// Symbol between two adjacent elements
	Sep string
	// Symbol at the start of the list
	Start string
}

const (
	// Default symbol at the end of the list
	DefaultEnd string = "]"
	// Default symbol between two adjacent elements of a doubly linked list
	DefaultSepDouble string = " <> "
	// Default symbol between two adjacent elements of a singly linked list
	DefaultSepSingle string = " > "
	// Default symbol at the start of the list
	DefaultStart string = "["
)

// Constructor for FormatSymbols
func NewFormatSymbols(isSingleLink bool) *FormatSymbols {
	var sep string

	if isSingleLink {
		sep = DefaultSepSingle
	} else {
		sep = DefaultSepDouble
	}

	return &FormatSymbols{
		End:   DefaultEnd,
		Sep:   sep,
		Start: DefaultStart,
	}
}

// Converts a singly or doubly linked list to a string
// according to format options and symbols
func formatToString[T any](
	formatOptions *at.Options, node commonNode[T], symbols *FormatSymbols,
) string {
	var builder strings.Builder
	builder.WriteString(symbols.Start)
	builder.WriteString(at.AnyToStringCustom(node.val(), formatOptions))

	for node.hasNext() {
		node = node.next()
		builder.WriteString(symbols.Sep)
		builder.WriteString(at.AnyToStringCustom(node.val(), formatOptions))
	}

	builder.WriteString(symbols.End)
	return builder.String()
}
