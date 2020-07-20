package problem0322

func CoinChange(coins []int, amount int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	// base case
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
