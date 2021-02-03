package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}
	left := fast
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Val, left.Val = left.Val, slow.Val
	return head
}

func main() {
	array := []int{100, 90}
	head := &ListNode{Val: array[0]}
	node := head
	for i := 1; i < len(array); i++ {
		node.Next = &ListNode{Val: array[i]}
		node = node.Next
	}

	newNode := swapNodes(head, 2)
	for newNode != nil {
		fmt.Println(newNode.Val)
		newNode = newNode.Next
	}
}
