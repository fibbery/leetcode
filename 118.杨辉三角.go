/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	rs := make([][]int, numRows)
	for i := range rs {
		rs[i] = make([]int, i+1)
		rs[i][0] = 1
		rs[i][i] = 1
		for j := 1; j < i; j++ {
			rs[i][j] = rs[i-1][j] + rs[i-1][j-1]
		}
	}
	return rs
}

// @lc code=end

