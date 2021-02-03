/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 *
 * https://leetcode-cn.com/problems/maximum-product-of-three-numbers/description/
 *
 * algorithms
 * Easy (50.07%)
 * Total Accepted:    45.9K
 * Total Submissions: 89.3K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
 *
 * 示例 1:
 *
 *
 * 输入: [1,2,3]
 * 输出: 6
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1,2,3,4]
 * 输出: 24
 *
 *
 * 注意:
 *
 *
 * 给定的整型数组长度范围是[3,10^4]，数组中所有的元素范围是[-1000, 1000]。
 * 输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
 *
 *
 */
package main

import "github.com/ethereum/go-ethereum/common/math"

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v < min1 {
			min1, min2 = v, min1
		} else if v < min2 {
			min2 = v
		}

		if v > max1 {
			max1, max2, max3 = v, max1, max2
		} else if v > max2 {
			max2, max3 = v, max2
		} else if v > max3 {
			max3 = v
		}

	}
	sum1 := max1 * max2 * max3
	sum2 := max1 * min1 * min2
	if sum1 > sum2 {
		return sum1
	} else {
		return sum2
	}
}

func main() {}
