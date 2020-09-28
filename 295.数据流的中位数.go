/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */
package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	data := []int{6, 10, 2, 6, 5, 0, 6, 3, 1, 0, 0}
	for i := 0; i < len(data); i++ {
		obj.AddNum(data[i])
		fmt.Println(obj.FindMedian())
	}
}

// @lc code=start

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
	if len(this.Max) >= len(this.Min) {
		//put into min
		if len(this.Max) != 0 && num > this.Max[0] {
			pop := this.Max[0]
			this.Max[0] = num
			this.Min = append(this.Min, pop)
			//最小值发生了改变，进行下沉操作，让二叉堆稳定
			sink(this.Max, lesser)
		} else {
			this.Min = append(this.Min, num)
		}
		// 插入末位，上浮改变二叉堆结构
		swim(this.Min, bigger)
	} else {
		//put into max
		if len(this.Min) != 0 && num < this.Min[0] {
			pop := this.Min[0]
			this.Min[0] = num
			this.Max = append(this.Max, pop)
			// 最大值发生了改变，下沉让二叉堆保持稳定
			sink(this.Min, bigger)
		} else {
			this.Max = append(this.Max, num)
		}
		// 插入末位， 上浮改变二叉堆结构
		swim(this.Max, lesser)
	}
}

func bigger(nums []int, i, j int) bool {
	if nums[i] >= nums[j] {
		return true
	} else {
		return false
	}
}

func lesser(nums []int, i, j int) bool {
	if nums[i] <= nums[j] {
		return true
	} else {
		return false
	}
}

func sink(nums []int, compareFn func([]int, int, int) bool) {
	length := len(nums)
	idx := 0
	for idx < length {
		left := 2*idx + 1
		right := 2*idx + 2
		if left > length-1 {
			return
		}
		tmp := left
		if right < length && compareFn(nums, right, left) {
			tmp = right
		}
		if compareFn(nums, idx, tmp) {
			break
		}
		nums[idx], nums[tmp] = nums[tmp], nums[idx]
		idx = tmp
	}
}

func swim(nums []int, compareFn func([]int, int, int) bool) {
	idx := len(nums) - 1
	for idx > 0 && compareFn(nums, idx, parentIdx(idx)) {
		nums[idx], nums[parentIdx(idx)] = nums[parentIdx(idx)], nums[idx]
		idx = parentIdx(idx)
	}
}

func parentIdx(idx int) int {
	return (idx - 1) >> 1
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Max) > len(this.Min) {
		return float64(this.Max[0])
	} else if len(this.Max) < len(this.Min) {
		return float64(this.Min[0])
	} else {
		return float64(this.Min[0]+this.Max[0]) / float64(2)
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
