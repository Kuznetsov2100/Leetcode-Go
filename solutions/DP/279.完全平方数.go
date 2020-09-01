/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	// 思路和leetcode-322零钱兑换相似，本题的完全平方数
	// 就是零钱兑换里的面值不同的硬币
	// dp[i]的含义是最少需要dp[i]个完全平方数的和凑出i
	dp := make([]int, n+1)
	dp[0] = 0 // base case
	for i := 1; i <= n; i++ {
		dp[i] = i // 一个正整数n最多需要n个1构成
		for j := 1; j*j <= i; j++ { // j*j <= i，以确保状态dp[i-j*j]是合法的
			dp[i] = min(dp[i], dp[i-j*j]+1) // 状态转移方程
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

