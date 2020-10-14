/*
 * @lc app=leetcode.cn id=1025 lang=golang
 *
 * [1025] 除数博弈
 */

// @lc code=start
func divisorGame(N int) bool {
	return N % 2 == 0
}

// 博弈型动态规划
// func divisorGame(N int) bool {
// 	dp := make([]int, N+1)
// 	dp[1] = false
// 	dp[2] = true
// 	for i := 3; i <= N; i++ {
// 		for j := 1; j < i; j++ {
// 			if i % j == 0 && !dp[i-j] {
// 				dp[i] = true
// 				break
// 			}
// 		}
// 	}
// 	return dp[N]
// }
// @lc code=end

