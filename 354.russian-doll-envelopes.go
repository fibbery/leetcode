/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode-cn.com/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (37.98%)
 * Total Accepted:    32.2K
 * Total Submissions: 80.5K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h)
 * 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 *
 * 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 *
 * 说明:
 * 不允许旋转信封。
 *
 * 示例:
 *
 * 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出: 3
 * 解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 *
 *
 */
package main

import "sort"

func maxEnvelopes(envelopes [][]int) (ret int) {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})
	dp := make([]int, len(envelopes)+1)
	for i := 1; i < len(dp); i++ {
		for j := 1; j < i; j++ {
			if envelopes[i-1][1] > envelopes[j-1][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ret = max(ret, dp[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
