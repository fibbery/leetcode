/*
 * @lc app=leetcode.cn id=1046 lang=golang
 *
 * [1046] 最后一块石头的重量
 */

// @lc code=start

type Heap struct {
	data []int
}

func NewHeap(values []int) *Heap {
	if len(values) <= 0 {
		return &Heap{data: make([]int, 0)}
	}
	heap := &Heap{data: values}
	beginIndex := len(values)/2 - 1
	for i := beginIndex; i >= 0; i-- {
		heap.sink(i)
	}
	return heap
}

func (h *Heap) Pop() int {
	if len(h.data) <= 0 {
		return -1
	}
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	ret := h.data[len(h.data)-1]
	h.data = h.data[0 : len(h.data)-1]
	h.sink(0)
	return ret
}
func (h *Heap) Push(value int) {
	h.data = append(h.data, value)
	h.swim(len(h.data) - 1)
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) sink(index int) {
	left := 2*index + 1
	right := 2*index + 2
	if left >= len(h.data) {
		return
	}
	largest := left
	if right < len(h.data) && h.data[right] > h.data[left] {
		largest = right
	}
	if h.data[largest] > h.data[index] {
		h.data[largest], h.data[index] = h.data[index], h.data[largest]
		h.sink(largest)
	}
}
func (h *Heap) swim(index int) {
	if index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] < h.data[index] {
			h.data[parent], h.data[index] = h.data[index], h.data[parent]
			h.swim(parent)
		}
	}
}

func lastStoneWeight(stones []int) int {
	heap := NewHeap(stones)
	for heap.Size() > 1 {
		x := heap.Pop()
		y := heap.Pop()
		if x > y {
			heap.Push(x - y)
		}
	}

	if heap.Size() == 1 {
		return heap.Pop()
	}
	return 0
}

// @lc code=end

