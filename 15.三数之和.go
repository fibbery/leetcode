/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) (ret [][]int) {
	if len(nums) < 3 {
		return
	}
	sort.Ints(nums)
	size := len(nums)
	for i := 0; i < size-2; i++ {
		//如果该值大于0，那么不可能使得三者大于0
		if nums[i] > 0 {
			return
		}
		// 如果值重复了，跳过循环继续，避免重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := size - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					//去除重复值
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					//去除右侧重复值
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return
}

// @lc code=end

