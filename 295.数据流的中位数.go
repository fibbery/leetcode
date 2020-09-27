/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
package main

type MedianFinder struct {
	// 一批大数，其实存储是一个小顶堆
	Max []int
	// 一批小数，其实存储是一个大顶堆
	Min []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{Max: make([]int, 0), Min: make([]int, 0)}
}

func (this *MedianFinder) AddNum(num int) {
}

func (this *MedianFinder) FindMedian() float64 {

}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
