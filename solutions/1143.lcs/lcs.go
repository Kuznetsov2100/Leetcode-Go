package problem1143

func LongestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// base case
	// dp[0][n] = 0
	// dp[m][0] = 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = func(x, y int) int { // func max(x,y int) int
					if x > y {
						return x
					}
					return y
				}(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
