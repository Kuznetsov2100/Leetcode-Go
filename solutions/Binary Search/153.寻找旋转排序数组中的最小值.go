/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right { // 搜索区间左闭右开
		mid := left + (right - left)/2
		if nums[mid] < nums[right] {
			right = mid // 最小值在靠近左半边，收缩右边界，nums[mid]可能是最小值，所以right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1 // 最小值靠近右半边，收缩左边界
		}
	}
	return nums[left] // left == right 
}
// @lc code=end

