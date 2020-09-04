/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	// 深度优先搜索+备忘录
	m := make(map[[2]int]int) // array in golang can be compared, so array can used as key in map
	var dfs func(i, substraction int) int
	dfs = func(i, substraction int) int {
		if i == len(nums) {
			if substraction == 0 {
				return 1
			}
			return 0
		}
		if val, ok := m[[2]int{i, substraction}]; ok {
			return val
		}
		ans := dfs(i+1, substraction+nums[i]) + dfs(i+1, substraction-nums[i])
		m[[2]int{i, substraction}] = ans
		return ans
	}
	return dfs(0, S)
}
// @lc code=end

