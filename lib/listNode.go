package lib

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(node *ListNode) {
	for {
		if node == nil {
			break
		}
		fmt.Println(node.Val)
		node = node.Next
	}
}
