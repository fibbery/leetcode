/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums) - 1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pos := partition(nums, left, right)
		quickSort(nums, left, pos - 1)
		quickSort(nums, pos + 1, right)
	}
}

func partition(nums []int, left, right int)int {
	pivot := left
	//从基准值后一个序号开始处理
	index := pivot + 1
	for j := index ; j <= right; j ++ {
		if nums[j] <= nums[pivot] {
			nums[j], nums[index] = nums[index], nums[j]
			index++
		}
	}
	nums[pivot] , nums[index - 1] = nums[index - 1] , nums[pivot]
    return index - 1
}


// @lc code=end

