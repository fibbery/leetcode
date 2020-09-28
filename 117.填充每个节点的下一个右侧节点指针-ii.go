/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
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

type LevelNode struct {
	*Node
	Level int
}

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]interface{}, 0),
	}
}

func (q *Queue) Push(value interface{}) {
	q.data = append(q.data, value)
}

func (q *Queue) Pop() interface{} {
	if len(q.data) < 1 {
		return nil
	}
	data := q.data[0]
	q.data = q.data[1:]
	return data
}

func (q *Queue) Size() int {
	return len(q.data)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := NewQueue()
	queue.Push(&LevelNode{
		Node:  root,
		Level: 0,
	})
	nodes := make([]*LevelNode, 0)
	for queue.Size() > 0 {
		node := queue.Pop().(*LevelNode)
		nodes = append(nodes, node)
		if node.Left != nil {
			queue.Push(&LevelNode{
				Node:  node.Left,
				Level: node.Level + 1,
			})
		}
		if node.Right != nil {
			queue.Push(&LevelNode{
				Node:  node.Right,
				Level: node.Level + 1,
			})
		}
	}
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i].Level == nodes[i+1].Level {
			nodes[i].Node.Next = nodes[i+1].Node
		}
	}
	return root
}

// @lc code=end

