/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] > nums[mid+1] { // 峰值在左侧，nums[mid]可能是峰值,所以right=mid
			right = mid
		} else if nums[mid] < nums[mid+1] { // 峰值在右侧，nums[mid]不可能是峰值，所以left = mid + 1
			left = mid + 1
		}
	}
	return left
}
// @lc code=end

