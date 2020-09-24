/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// package main

// type Node struct {
// 	Val               int
// 	Left, Right, Next *Node
// }

type LevelNode struct {
	*Node
	Level int
}

func connect(root *Node) *Node {
	queue := NewQueue()
	queue.Offer(&LevelNode{Node: root, Level: 0})
	nodes := make([]*LevelNode, 0)
	for queue.Len() > 0 {
		node := queue.Pop()
		nodes = append(nodes, node)
		if node.Left != nil {
			queue.Offer(&LevelNode{
				Node:  node.Left,
				Level: node.Level + 1,
			})
		}
		if node.Right != nil {
			queue.Offer(&LevelNode{
				Node:  node.Right,
				Level: node.Level + 1,
			})
		}
	}
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Right != nil {
			nodes[i].Left.Next = nodes[i].Right
		} else {
			continue
		}

		if i+1 <= i+nodes[i].Level {
			nodes[i].Right.Next = nodes[i+1].Left
		}
	}
	return root
}

type Queue struct {
	data []*LevelNode
}

func NewQueue() *Queue {
	return &Queue{data: make([]*LevelNode, 0)}
}

func (q *Queue) Offer(node *LevelNode) {
	q.data = append(q.data, node)
}

func (q *Queue) Poll() *LevelNode {
	return q.data[0]
}

func (q *Queue) Pop() *LevelNode {
	pop := q.data[0]
	q.data = q.data[1:]
	return pop
}

func (q *Queue) Len() int {
	return len(q.data)
}

// @lc code=end
