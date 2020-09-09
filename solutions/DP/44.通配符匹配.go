/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// dp[0][0]=True，即当字符串s和模式p均为空时，匹配成功；
	// dp[i][0]=False，即空模式无法匹配非空字符串；
	// dp[0][j] 需要分情况讨论：因为星号才能匹配空字符串，所以只有当模式p的前j个字符均为星号时，dp[0][j]才为真。
	for j := 1; j <= n; j++ {
        if p[j-1] != '*' {
			break
		}
		dp[0][j] = true
    }
	
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' { 
			// dp[i-1][j]表示使用了这个星号, 比如： s: abb p:a* 
			// dp[i][j-1]表示没有使用这个星号, 比如: s: abb p:abb*
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
// @lc code=end

 