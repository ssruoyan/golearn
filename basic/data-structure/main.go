package main

import (
	"./linkedlist.go"
	"fmt"
)

func main() {
	list := LinkedList{}
	l1 := Node{3, nil}

	list.Append(&l1)

	fmt.Println("Data: ", list)
}