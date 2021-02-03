/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return merge(sort(head, slow), sort(slow, tail))
}

func merge(n1, n2 *ListNode) *ListNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	head := &ListNode{}
	dummy := head
	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			dummy.Next = n1
			n1 = n1.Next
		} else {
			dummy.Next = n2
			n2 = n2.Next
		}
		dummy = dummy.Next
	}

	if n1 == nil {
		dummy.Next = n2
	}
	if n2 == nil {
		dummy.Next = n1
	}
	return head.Next
}

// @lc code=end

func main() {
	array := []int{9, 9, 8, 10, 25, 1}
	head := &ListNode{Val: array[0]}
	node := head
	for i := 1; i < len(array); i++ {
		node.Next = &ListNode{Val: array[i]}
		node = node.Next
	}
	n := sortList(head)
	for ; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}

}
