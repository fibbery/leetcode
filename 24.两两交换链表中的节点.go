/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverse(head, 2)
	head.Next = swapPairs(head.Next)
	return newHead
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	a, b := head, head.Next
	a.Next = swapPairs2(b.Next)
	b.Next = a
	return b
}

var successor *ListNode

func reverse(head *ListNode, k int) *ListNode {
	if k == 1 {
		successor = head.Next
		return head
	}
	if head == nil {
		return nil
	}
	newHead := reverse(head.Next, k-1)
	head.Next.Next = head
	head.Next = successor
	return newHead
}

// @lc code=end

func main() {
	array := []int{1, 2, 3, 4}
	head := &ListNode{Val: array[0]}
	node := head
	for i := 1; i < len(array); i++ {
		node.Next = &ListNode{Val: array[i]}
		node = node.Next
	}
	n := swapPairs3(head)
	for ; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
}
