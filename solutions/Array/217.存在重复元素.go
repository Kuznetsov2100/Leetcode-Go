/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}
// @lc code=end

