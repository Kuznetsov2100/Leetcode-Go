/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 */

// @lc code=start
func singleNonDuplicate(nums []int) int {
	// 经分析，这个单一元素只可能出现在数组中偶数索引的位置上，所以可以对偶数索引执行二分搜索
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// 如果mid为奇数，-1使得mid为偶数
		if mid % 2 == 1 {
			mid--
		}
		// 如果nums[mid] == nums[mid+1]，那么[left,mid+1]区间为偶数个数字，
		// 不可能存在单一元素，所以缩小左边界
		if nums[mid] == nums[mid+1] {
			left = mid + 2 
		// 如果nums[mid] != nums[mid+1]，因为[mid+1, right]区间为偶数个数字，
		// 不可能存在单一元素， 所以缩小有边界
		} else {
			right = mid
		}
	}
	return nums[left] // 循环结束后，left等于right，搜索区间只剩一个元素，这就是数组的单一元素
}
// @lc code=end

