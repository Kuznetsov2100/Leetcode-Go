/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	// dp[i][0]: in day i, not hold stock
	// dp[i][1]: in day i, hold stock
	// dp[n-1][0]: in last day, we got the maximum profit after stock has been sold.
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	// base case
	// dp[-1][0] = 0
	// dp[-1][1] = -infinity
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			// dp[i][0] = max(dp[-1][0], dp[-1][1]+prices[0])
			// dp[i][0] = max(0, -infinity + prices[0]) = 0
			dp[i][1] = -prices[0]
			// dp[i][1] = max(dp[-1][1],dp[-1][0]-prices[0])
			// dp[i][1] = max(-infinity, 0-prices[0]) = -prices[0]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// dp[i][0] = max(choose rest, choose sell)
		dp[i][1] = max(dp[i-1][1], -prices[i]) // dp[i-1][0] = 0
		// dp[i][1] = max(choose rest, choose buy)
	}
	return dp[n-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

