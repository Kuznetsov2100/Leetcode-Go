/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] < nums[mid] { // 左边有序
			if nums[left] <= target && target <= nums[mid] { // 如果target处于左边范围，则在左边查找
				right = mid-1 // 收缩右侧边界
			} else {
				left = mid + 1 // target不在左边，收缩左侧边界
			}
		} else if nums[left] > nums[mid] { // 右边有序
			if nums[mid] <= target && target <= nums[right] { // 如果target处于右边范围，则在右边查找
				left = mid + 1 // 收缩左侧边界
			} else {
				right = mid - 1 // target不在右边，收缩右侧边界
			}
		} else if nums[left] == nums[mid] { // 无法判断左边或者右边是否有序，递增left，跳过与nums[mid]相等的值
			left++
		}
	}
	return false
}
// @lc code=end

