/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 利用153题的思路找到数组最小值的索引
	for left < right { // 搜索区间左闭右开
		mid := left + (right - left)/2
		if nums[mid] < nums[right] {
			right = mid // 最小值在靠近左半边，收缩右边界，nums[mid]可能是最小值，所以right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1 // 最小值靠近右半边，收缩左边界
		}
	}
	if nums[left] == target {
		return left // left为数组最小值的索引， 如果target为最小值，直接返回
	} else if target > nums[left] && target <= nums[len(nums)-1] {
		return helper(nums, left+1,len(nums)-1, target)	// 如果target在最小值右边的区间(有序),执行二分搜索	
	} else {
		return helper(nums, 0, left-1, target) // 如果target在最小值左边的区间(有序),执行二分搜索
	}
}

func helper(nums []int, start, end, target int) int {
	left, right := start, end
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
// @lc code=end

