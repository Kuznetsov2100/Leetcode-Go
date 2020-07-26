/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	n := len(nums)
	res := 0
	// dp[i] 表示以 nums[i] 这个数结尾的最长递增子序列的长度
	dp := make([]int,n)
	for i := 0; i < n; i++ {
		dp[i] = 1 // 子序列数组最小长度为1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i],dp[j]+1)
			}
		}
		res = max(res,dp[i])
	}
	return res
}
	
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

