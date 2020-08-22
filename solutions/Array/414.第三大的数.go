/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	n := len(nums)
	sort.Ints(nums) // 首先对数组排序
	if n < 3 { // 如果数组长度小于3，直接返回排序后数组最后一个数字
		return nums[n-1]
	}
	// 消除数组重复元素
	slow, fast := 0, 1 // 快慢指针
	// 我们让慢指针 slow 走左后面，快指针 fast 走在前面探路，找到一个不重复的元素就告诉 slow 并让 slow 前进一步。
	// 这样当 fast 指针遍历完整个数组 nums 后，nums[0..slow] 就是不重复元素，之后的所有元素都是重复元素。
	for fast < n {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	newlength := slow+1 // 无重复元素的数组长度
	if newlength < 3 {
		return nums[n-1] 
	}
	return nums[newlength-3]
}
// @lc code=end

