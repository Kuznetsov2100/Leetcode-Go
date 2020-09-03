/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// 状态压缩到一维数组
	dp := make([]int, n) // 上一行的状态
	for i := range dp {
		dp[i] = 1 // 初始化, 第一行全为1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1] // 当前行位置j的状态由左边和上一行位置j的状态决定
		}
	}
	return dp[n-1] // 最后一行的最后一个位置即为结果
}

// 动态规划: 二维数组
// func uniquePaths(m int, n int) int {
// 	dp := make([][]int, m)
// 	for i := range dp {
// 		dp[i] = make([]int, n)
// 	}
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if i == 0 && j == 0 {
// 				dp[i][j] = 1
// 				continue
// 			}
// 			if i == 0 {
// 				dp[i][j] = dp[i][j-1]
// 				continue
// 			}
// 			if j == 0 {
// 				dp[i][j] = dp[i-1][j]
// 				continue
// 			}
// 			dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 		}
// 	}
// 	return dp[m-1][n-1]
// }
// @lc code=end

