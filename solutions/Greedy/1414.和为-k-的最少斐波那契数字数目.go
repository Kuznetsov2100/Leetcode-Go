/*
 * @lc app=leetcode.cn id=1414 lang=golang
 *
 * [1414] 和为 K 的最少斐波那契数字数目
 */

// @lc code=start
func findMinFibonacciNumbers(k int) int {
	// 贪心思想， 小于等于10^9的斐波那契数一共有44个
	// 倒序遍历斐波那契数列，每次减去小于等于k的最大斐波那契数,直到k等于0
	n := 44
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n-1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	var res int
	for i := n-1; i >= 0; i-- {
		if k == 0 {
			break
		}
		if k >= dp[i] {
			k -= dp[i]
			res++
		}
	}
	return res
}
// @lc code=end

