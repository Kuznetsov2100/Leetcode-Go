/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	res := -1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			res = mid
			break
		}
	}
	if res != -1 {
		return res
	}
	if left == len(nums) {
		return left
	}
	if right < 0 {
		return right+1
	}
	return left
}
// @lc code=end

