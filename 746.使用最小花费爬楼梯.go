/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	prev, current := 0, 0
	for i := 2; i < len(dp); i++ {
		prev, current = current, min(prev+cost[i-2], current+cost[i-1])
	}
	return current
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

