/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
// 本题是经典0-1背包问题的变体，选一个物品需要消耗两种价值(字符0的容量，字符1的容量)
// 版本一的代码使用三维dp[i][j][k]表示对于前i个字符串,使用了j个0和k个1能够拼出的字符串的最大数量
// 因为dp[i][j][k]只与dp[i-1][j][k], dp[i-1][j-count0][k-count1]有关，可以进行空间优化，
// 将状态压缩到二维：dp[i][j] = max(dp[i][j], dp[i-count0][j-count1]+1),
// 右边的dp[i][j]即是dp[i-1][j][k]， dp[i-count0][j-count1]+1 等价于 dp[i-1][j-count0][k-count1]+1
// 前向遍历时， dp[i][j]依赖其左上方的上一轮的旧值，但是在遍历到dp[i][j]时，其左上方的旧值已被新值覆盖，
// 因此dp[i][j]存储的是错误的值。
// 采用倒序遍历才可以正确计算状态：
// 倒序遍历每次错右下角开始，dp[i][j]左上方存储的是上一轮的旧值，在左上角的值被新值覆盖时，右下角的值已得到正确更新


// 版本二
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		count0 := strings.Count(str, "0")
		count1 := len(str)-count0
		for i := m; i >= count0; i-- {
			for j := n; j >= count1; j-- {
				dp[i][j] = max(dp[i][j], dp[i-count0][j-count1]+1)
			}
		}
	}
	return dp[m][n]	
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 版本一
// func findMaxForm(strs []string, m int, n int) int {
// 	dp := make([][][]int, len(strs)+1)
// 	for i := range dp {
// 		dp[i] = make([][]int, m+1)
// 		for j := range dp[i] {
// 			dp[i][j] = make([]int, n+1)
// 		}
// 	}

// 	for i := 1; i <= len(strs);i++ {
// 		count0 := strings.Count(strs[i-1], "0")
// 		count1 := len(strs[i-1])-count0
// 		for j := 0; j <= m; j++ {
// 			for k := 0; k <= n; k++ {
// 				if j - count0 >= 0 && k-count1 >= 0 {
// 					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-count0][k-count1]+1)
// 				} else {
// 					dp[i][j][k] = dp[i-1][j][k]
// 				}
// 			}
// 		}
// 	}
// 	return dp[len(strs)][m][n]
// }
// @lc code=end

