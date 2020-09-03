/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
// 直接使用原数组存储状态
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
				continue
			}
			if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
				continue
			}
			if j == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
				continue
			}
			obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
		}
	}
	return obstacleGrid[m-1][n-1]
}

// 动态规划: 二维数组
// func uniquePathsWithObstacles(obstacleGrid [][]int) int {
// 	m, n := len(obstacleGrid), len(obstacleGrid[0])
// 	dp := make([][]int, m)
// 	for i := range dp {
// 		dp[i] = make([]int, n)
// 	}
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if obstacleGrid[i][j] == 1 {
// 				continue
// 			}
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

