/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */
package main

// @lc code=start
type LRUCache struct {
	Capacity int
	M        map[int]*Node
	L        *LinkedList
}

type Node struct {
	key, value int
	prev, next *Node
}

func NewNode(key, val int) *Node {
	return &Node{key: key, value: val}
}

type LinkedList struct {
	head, tail *Node
	size       int
}

func NewLinkedList() *LinkedList {
	head := new(Node)
	tail := new(Node)
	head.next = tail
	tail.prev = head
	return &LinkedList{head: head, tail: tail}
}

func (l *LinkedList) AddFirst(node *Node) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
	l.size++
}

func (l *LinkedList) RemoveLast() *Node {
	prev := l.tail.prev
	if prev != l.head {
		prev.prev.next = l.tail
		l.tail.prev = prev.prev
		l.size--
		return prev
	}
	return nil
}

func (l *LinkedList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
	l.size--
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		M:        make(map[int]*Node),
		L:        NewLinkedList(),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.M[key]; ok {
		this.L.Remove(v)
		this.L.AddFirst(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key, value int) {
	node := NewNode(key, value)
	if v, ok := this.M[key]; ok {
		this.L.Remove(v)
		this.L.AddFirst(node)
		this.M[key] = node
	} else {
		if this.L.size >= this.Capacity {
			last := this.L.RemoveLast()
			if last != nil {
				delete(this.M, last.key)
			}
		}
		this.L.AddFirst(node)
		this.M[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
