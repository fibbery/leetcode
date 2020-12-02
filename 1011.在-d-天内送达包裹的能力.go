/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
func shipWithinDays(weights []int, D int) int {
	//最低运载能力：重量最大值，保证每天能运载一次
	//最高运载能力：重量总和，保证一次运完
	left, right := getMaxWeight(weights), getSumWeight(weights)
	//有效性校验
	if left == right {
		return -1
	}

	for left <= right {
		mid := left + (right-left)/2
		if canShip(weights, mid, D) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func getMaxWeight(weights []int) int {
	if len(weights) <= 0 {
		return -1
	}
	max := weights[0]
	for _, v := range weights {
		if v > max {
			max = v
		}
	}
	return max
}

func getSumWeight(weights []int) int {
	if len(weights) <= 0 {
		return -1
	}
	sum := 0
	for _, v := range weights {
		sum += v
	}
	return sum
}

func canShip(weights []int, cap, days int) bool {
	remain := days
	for remain > 0 {
		i := 0
		c := cap
		for i < len(weights) && c >= weights[i] {
			c = c - weights[i]
			i++
		}
		weights = weights[i:]
		remain--
	}
	return len(weights) == 0
}

// @lc code=end

