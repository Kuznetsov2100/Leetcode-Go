/*
 * @lc app=leetcode.cn id=1403 lang=golang
 *
 * [1403] 非递增顺序的最小子序列
 */

// @lc code=start
func minSubsequence(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}
	sort.Ints(nums)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}
	var res []int
	for i := 1; i <= n; i++ {
		res = append(res, nums[n-i])
		if dp[n]-dp[n-i] > dp[n-i] {
			break
		}
	}
	return res
}
// @lc code=end

