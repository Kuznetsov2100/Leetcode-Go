/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	var res []int
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			for nums[i] != nums[nums[i]-1] {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			res = append(res, nums[i])
		}
	}
	return res
}
// @lc code=end

