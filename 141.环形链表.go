package main

import "fmt"

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	visted := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := visted[head]; ok {
			return true
		}
		visted[head] = struct{}{}
		head = head.Next
	}
	return false
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 4}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b
	fmt.Println(hasCycle(a))
}
