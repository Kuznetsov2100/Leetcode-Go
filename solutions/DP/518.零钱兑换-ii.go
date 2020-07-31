/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	n := len(coins)
	// dp[i][j]的含义:若只使用coins中的前i个硬币的面值,若想凑出金额 j,有dp[i][j]种凑法
	dp := make([][]int,n+1)
	for i := range dp {
		dp[i] = make([]int,amount+1)
		dp[i][0] = 1 // base case 凑出金额0有1种凑法
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j - coins[i-1] >= 0 {
				// 假如不使用coins[i-1]面值的硬币，凑出金额j的凑法继承自dp[i-1][j]
				// 加入使用coins[i-1]面值的硬币，凑出金额j的凑法应关注如何出金额j-coins[i-1]
				// 共有多少种凑法，所以dp[i][j]的值为上述两种选择的和。
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][amount]
}
// @lc code=end

