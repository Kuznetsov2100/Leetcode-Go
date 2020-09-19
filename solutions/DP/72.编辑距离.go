/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	// dp[i][j]表示word1前i个字符转换到word2前j个字符所使用的最少操作数
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// word2为空，删除所有word1字符即可转换到word2
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	// word1为空，插入word2中的每一个字符即可转换到word2
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 无需操作，继承前面的状态
			} else {
				// dp[i-1][j]+1, 表示删除word1的一个字符
				// dp[i][j+1]+1, 表示添加一个字符到word1中
				// dp[i-1][j-1]+1, 表示修改word1中的一个字符
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

