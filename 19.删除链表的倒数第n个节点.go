/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */
package main

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := dummy, head
	for i := 0; i < n; i++ {
		second = second.Next
	}
	for ; second != nil; second = second.Next {
		first.Next = first
	}
	first.Next = first.Next.Next
	return dummy.Next
}

// @lc code=end
