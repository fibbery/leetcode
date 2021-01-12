/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 * 思路：利用map，用值作为map的key，索引作为map的value。正好可以推断出对应两数之和的索引
 */
package main

import "fmt"

// @lc code=start
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for index, value := range nums {
		key := target - value
		_, ok := dict[key]
		if ok {
			return []int{index, dict[key]}
		} else {
			dict[value] = index
		}
	}
	return make([]int, 0)
}

// @lc code=end

func main() {
	fmt.Println("hello world")
}
