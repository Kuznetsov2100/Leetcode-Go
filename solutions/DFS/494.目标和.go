/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	var res int
	var dfs func(i, track int)
	dfs = func(i, track int) {
		if i == len(nums) && track == 0 {
			res++
			return
		}
		if i == len(nums) {
			return
		}
		track += nums[i]
		dfs(i+1, track)
		track -= nums[i]

		track -= nums[i]
		dfs(i+1, track)
		track += nums[i]
	}
	dfs(0, S)
	return res
}
// @lc code=end

