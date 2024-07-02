package main

import (
	"fmt"

	lk "github.com/Matej-Chmel/go-linked-list"
)

func main() {
	singleHead := lk.CreateSinglyLinkedList(1, 2, 3)
	doubleHead := lk.CreateDoublyLinkedList[int8](10, 20, 30)

	fmt.Println(singleHead)
	fmt.Println(doubleHead)
}
