/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1) // dp[i]的含义是将正整数i拆分成至少两个正整数的和之后，这些正整数的最大乘积
	dp[1] = 0 // 1无法拆分成两个正整数的和
	// 当i >=2时，假设对正整数i拆分出的第一个正整数是j(1 <= j < i),则有以下两种情况：
	// case1: 将i拆分成j和i-j的和，此时的乘积是j*(i-j)
	// case2: 将i拆分成j和i-j的和，其中i-j是多个正整数的和，此时的乘积是j*dp[i-j]
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

