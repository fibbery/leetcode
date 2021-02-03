/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// Package main provides ...
package main

// @lc code=start
func prefixesDivBy5(A []int) (ret []bool) {
	x := 0
	for i := 0; i < len(A); i++ {
		x = (x<<1 | A[i]) % 5
		ret = append(ret, x == 0)
	}
	return
}

// @lc code=end
