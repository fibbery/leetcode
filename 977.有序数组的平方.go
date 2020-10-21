/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */
package main

// @lc code=start
func sortedSquares(A []int) []int {
	result := make([]int, len(A))
	i, j := 0, len(A)-1
	for x := len(A) - 1; x >= 0; x-- {
		if A[i]*A[i] > A[j]*A[j] {
			result[x] = A[i] * A[i]
			i++
		} else {
			result[x] = A[j] * A[j]
			j--
		}
	}
	return result
}

// @lc code=end
