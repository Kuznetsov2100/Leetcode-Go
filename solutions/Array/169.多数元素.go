/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	// Boyer-Moore 投票算法
	candidate, votes := nums[0], 1
	for _, v := range nums {
		if candidate == v {
			votes++
		} else {
			votes--
		}
		if votes == 0 {
			candidate = v
			votes = 1
		}
	}
	return candidate
}
// @lc code=end

