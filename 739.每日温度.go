/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
type Item struct {
	temperature int
	index       int
}

type ItemStack []*Item

func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))
	var stack ItemStack
	for index, value := range T {
		if len(stack) == 0 {
			stack = append(stack, &Item{value, index})
			continue
		}
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if value > top.temperature {
				result[top.index] = index - top.index
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, &Item{value, index})
	}
	return result
}
// @lc code=end

