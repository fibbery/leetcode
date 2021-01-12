/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	// 声明一个三维数组 dp[i][j][0 or 1]
	// 1 <= i <= len(prices) 1 <= j <= k
	// 0 or 1 表示当天是持有还是卖出
	// dp[1][1][0] 表示当天第一天，最多允许操作1次，当天不持有股票能获得的最大利润
	// dp[1][1][1] 表示当天第一天，最多允许操作1次，当天持有股票时能获得的最大利益(这里包括已有利润值+当天持有股票值）

	//初始化数组
	dp := make([][][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
			if i == 0 {
				//base case
				dp[0][j][1] = math.MinInt32
			}
		}
		// base case 如果允许操作0次，那么拥有股票的几率为0，这种情况不可能出现
		dp[i][0][1] = math.MinInt32
	}

	for i := 1; i <= len(prices); i++ {
		for j := 1; j <= k; j++ {
			//当天不持有股票，那就是昨天持有股票今天不操作，或者是昨天持有股票今天卖了
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			//当天持有股票，那就是昨天持有股票今天不操作，或者昨天无股票，今天买了, 注意j - 1, 因为卖出才算完整的一次交易
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}

	// 返回最后一天，不持有股票时的最大利润，这样才能称得上是k次交易
	return dp[len(prices)][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

