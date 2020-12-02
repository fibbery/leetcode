/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */

package main

import "fmt"

func main() {
	fmt.Println(maximumGap([]int{100, 3, 2, 1}))
}

// @lc code=start
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	max := maxNum(nums)
	min := minNum(nums)

	// 平均间距
	avg := maxNum([]int{1, (max - min) / (len(nums) - 1)})
	bucketSize := (max-min)/avg + 1
	buckets := make([]pair, bucketSize)
	// init buckets
	for i := 0; i < bucketSize; i++ {
		buckets[i].min, buckets[i].max = -1, -1
	}
	// 将每个数据落到对应的桶上面, 桶是按照 avg间距进行分割
	for i := 0; i < len(nums); i++ {
		d := (nums[i] - min) / avg
		if buckets[d].min == -1 {
			buckets[d].min = nums[i]
			buckets[d].max = nums[i]
		} else {
			buckets[d].min = minNum([]int{buckets[d].min, nums[i]})
			buckets[d].max = maxNum([]int{buckets[d].max, nums[i]})
		}
	}

	mostFar := 0
	// 前一个桶的index, 因为可能碰到空桶
	prev := -1
	for i := 0; i < bucketSize; i++ {
		currentBucket := buckets[i]
		if currentBucket.min == -1 {
			//桶没被初始化
			continue
		}
		if prev != -1 {
			prevBucket := buckets[prev]
			mostFar = maxNum([]int{mostFar, currentBucket.min - prevBucket.max})
		}
		prev = i
	}
	return mostFar
}

type pair struct {
	min, max int
}

func maxNum(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func minNum(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

// @lc code=end
