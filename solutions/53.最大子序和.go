/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	n := len(nums) // 题目条件中nums长度不为0,无需判断n
	// dp[i]的含义： 以nums[i]为结尾的最大子数组和为dp[i]。
	dpCurrentSum := nums[0] // base case
	res := nums[0]
	for i := 1; i < n; i++ {
		// 注意dp[i]的含义，要么以nums[i]结尾形成了一个更大的连续子数组，
		// 要么形成只有nums[i]一个元素的的子数组
		// 压缩空间复杂度，无需dp数组
		dpNewSum  := max(nums[i], nums[i]+ dpCurrentSum)
		dpCurrentSum = dpNewSum
		res = max(res,dpNewSum)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
func maxSubArray(nums []int) int {
	n := len(nums) // 题目条件中nums长度不为0,无需判断n
	// dp[i]的含义： 以nums[i]为结尾的最大子数组和为dp[i]。
	dp := make([]int,n)
	dp[0] = nums[0] // base case
	sum := nums[0]
	for i := 1; i < n; i++ {
		// 注意dp[i]的含义，要么以nums[i]结尾形成了一个更大的连续子数组，
		// 要么形成只有nums[i]一个元素的的子数组
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	// 遍历dp数组，找到最大的子数组和
	for i := range dp {
		sum = max(sum,dp[i])
	}
	return sum
}
*/

// @lc code=end

