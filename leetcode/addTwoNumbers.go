package main

import (
	"fmt"
)

func main() {
	l1 := toLinkedList([]int{2, 4, 3})

	fmt.Println("result: ", l1)
}


type LinkNode struct {
	Val int
	Next *LinkNode
}

func addTwoNumbers(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	// 记录进位
	carry := 0

	for l1 != nil || l2 != nil {
		carry = add(l1, l2, carry)

		if l1 != nil && l1.Next != nil {
			l1 = l1.Next
		}
		if l2 != nil && l2.Next != nil {
			l2 = l2.Next
		}
	}

	return l1
}

func add(l1 *LinkNode, l2 *LinkNode, carry int) int {
	l1.Val = (l1.Val + l2.Val + carry) % 10
	carry = (l1.Val + l2.Val + carry) / 10

	return carry
}

func toLinkedList(arr []int) []LinkNode{
	var list []LinkNode

	for _, x := range arr {
		var node LinkNode

		(node).Val = x
		(node).Next = nil

		list = append(list, node)
	}

	for i, n := range list {
		if i + 1 < len(list) {
			n.Next = &(list[i + 1])
		}
		fmt.Println(i, n)
	}

	return list
}