# Go Linked List
Generic library for singly and doubly linked lists.

## Installation
```bash
go get github.com/Matej-Chmel/go-linked-list@v1.0.1
```

## Features
- Easy to use
- Methods for accessing nodes by index and direction
- Customizable conversion to a string

## Example
```go
package main

import (
	"fmt"

	lk "github.com/Matej-Chmel/go-linked-list"
)

func main() {
	// Easy construction
	singleHead := lk.CreateSinglyLinkedList(1, 2, 3)
	doubleHead := lk.CreateDoublyLinkedList[int8](10, 20, 30)

	// Conversion to string
	fmt.Println(singleHead)
	fmt.Println(doubleHead)

	node := singleHead
	sum := 0

	for node != nil {
		// Each node has Next and Val fields
		sum += node.Val
		node = node.Next
	}

	fmt.Println("Sum:", sum)

	// Doubly linked list can be iterated backwards
	current := doubleHead.GetLast()
	total := int8(0)

	for current != nil {
		total += current.Val
		current = current.Prev
	}

	fmt.Println("Total:", total)

	// Conversion to string can be customized
	symbols := lk.NewFormatSymbols(true)
	symbols.Start = "("
	symbols.Sep = ", "
	symbols.End = ")"

	fmt.Println(singleHead.Format(symbols))
}
```

### Output
```none
[1 > 2 > 3]
[10 <> 20 <> 30]
Sum: 6
Total: 60
(1, 2, 3)
```
