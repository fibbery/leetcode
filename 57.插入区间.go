/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */
package main

import "fmt"

func main() {
	fmt.Printf("result is : %+v", insert([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{4, 8}))
}

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	isMerge := false
	var ans [][]int
	for _, interval := range intervals {
		if interval[0] > right {
			if !isMerge {
				ans = append(ans, []int{left, right})
				isMerge = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			ans = append(ans, interval)
		} else {
			left = min(interval[0], left)
			right = max(interval[1], right)
		}
	}

	if !isMerge {
		ans = append(ans, []int{left, right})
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
