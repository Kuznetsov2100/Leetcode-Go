/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	length := len(prices)
	maxDeals := 2
	if length <= 1 {
		return 0
	}
	// dp[i][k][0]: in day i, has made k deals, not hold stock
	// dp[i][k][1]: in day i, has made k deals, hold stock
	// dp[n-1][maxDeals][0]: in last day, has made maxDeals deals, we got the maximum profit after stock has been sold.
	dp := make([][][]int,length)
	for i := range dp {
		dp[i] = make([][]int, maxDeals+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	
	for today := 0; today < length; today++ {
		for deals := 1; deals <= maxDeals; deals++ {
			if today == 0  { // base case
				dp[today][deals][0] = 0
				dp[today][deals][1] = -prices[0]
				continue
			}
			dp[today][deals][0] = max(dp[today-1][deals][0], dp[today-1][deals][1]+ prices[today])
			dp[today][deals][1] = max(dp[today-1][deals][1], dp[today-1][deals-1][0]-prices[today])
		}
	}
	return dp[length-1][maxDeals][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

