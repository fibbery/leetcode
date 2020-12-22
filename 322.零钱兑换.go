/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	//初始化一个备忘率
	dptable := make([]int, amount+1)
	for i := 1; i < len(dptable); i++ {
		dptable[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dptable[i] = min(dptable[i], 1+dptable[i-coin])
		}
	}
	return dptable[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

