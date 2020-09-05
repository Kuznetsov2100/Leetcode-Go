/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start

func removeDuplicates(nums []int) int {
	// 定义nums[0,index)为没有重复元素的数组区间
	index := 0
	for i := 0; i < len(nums); i++ {
		if index == 0 || nums[i] != nums[index-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index // index-1 - 0 + 1
}

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	slow, fast := 0, 1 // 快慢指针
// 	// 我们让慢指针 slow 走左后面，快指针 fast 走在前面探路，找到一个不重复的元素就告诉 slow 并让 slow 前进一步。
// 	// 这样当 fast 指针遍历完整个数组 nums 后，nums[0..slow] 就是不重复元素，之后的所有元素都是重复元素。
// 	for fast < len(nums) {
// 		if nums[fast] != nums[slow] {
// 			slow++
// 			nums[slow] = nums[fast]
// 		}
// 		fast++
// 	}
// 	return slow+1
// }
// @lc code=end

