/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 */
package main

import "fmt"

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, 0)
	for index, value := range nums {
		count := 0
		for i := 0; i < len(nums); i++ {
			if i != index && value > nums[i] {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}

// @lc code=end
