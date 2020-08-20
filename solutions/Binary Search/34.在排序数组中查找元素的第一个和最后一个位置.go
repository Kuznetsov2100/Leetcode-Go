/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	return []int{bound(nums, target,0), bound(nums, target, 1)}
}

// sign为0用于搜索左侧第一个位置，sign为1用于搜索右侧最后一个位置
func bound(nums []int, target int, sign int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			if sign == 0 {
				right = mid - 1 // 搜索左侧边界时，锁定右侧边界
			} else {
				left = mid + 1 // 搜索右侧边界时，锁定左侧边界
			}
		}
	}
	if sign == 0 { // 搜索左侧边界时，检查left越界情况
		if left >= len(nums) || nums[left] != target {
			return -1
		}
		return left // target第一个位置
	} else { // 搜索右侧边界时，检查right越界情况
		if right < 0 || nums[right] != target {
			return -1
		}
		return right // target最后一个位置
	}
}
// @lc code=end

