/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	length := len(prices)
	maxDeals := k
	if length <= 1 || maxDeals < 1 {
		return 0
	}
	// 一次交易由买入和卖出构成，至少需要两天。所以说有效的限制 maxDeals 应该不超过 length/2，
	// 如果超过，就没有约束作用了，相当于 maxDeals = +infinity
	if maxDeals > length/2 {
		return maxProfitInf(prices)
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

// 没有交易次数限制，maxDeals = +infinity
func maxProfitInf(prices []int) int {  
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
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) 
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

