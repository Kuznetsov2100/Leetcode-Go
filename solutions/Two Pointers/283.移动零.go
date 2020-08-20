/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		} else if nums[slow] != 0 || nums[fast] != 0 {
			slow++
		}
		fast++
	}
}
// @lc code=end

