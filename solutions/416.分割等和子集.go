/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	// 规约为0-1背包问题模型
	// 给一个可装载重量为sum/2的背包和N个物品，每个物品的重量为nums[i]。
	// 现在让你装物品，是否存在一种装法，能够恰好将背包装满？
	n := len(nums)
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum % 2 != 0 { // 数组元素和为奇数，数组无法分割成等和子集
		return false
	}
	sum = sum / 2
	// dp[i][j] = x 表示，对于前 i 个物品，当前背包的容量为 j 时，若 x 为 true，
	// 则说明可以恰好将背包装满，若 x 为 false，则说明不能恰好将背包装满。
	dp := make([][]bool, n+1)
	// base case 就是 dp[..][0] = true 和 dp[0][..] = false，因为背包没有空间的时候，
	//就相当于装满了，而当没有物品可选择的时候，肯定没办法装满背包。
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true // base case
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j - nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}
// @lc code=end

