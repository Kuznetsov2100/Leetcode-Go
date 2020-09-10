/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	len1, len2, len3 := len(s1), len(s2), len(s3)
	if len1 + len2 != len3 {
		return false // s1和s2的长度之和不等于s3的长度，则s3一定无法由s1和s2交错组成
	}
	// dp[i][j] 表示s1的前i个字符和s2的前j个字符能否交错组成s3的前i+j个字符
	dp := make([][]bool, len1+1)
	for i := range dp {
		dp[i] = make([]bool, len2+1)
	}
	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			k := i+j-1
			if i == 0 && j == 0 {
				dp[i][j] = true // base case
				continue
			}
			if i == 0 { // 边界条件
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[k]
				continue
			}
			if j == 0 { // 边界条件
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[k]
				continue
			}		
			dp[i][j] = (dp[i][j-1] && s2[j-1] == s3[k]) || (dp[i-1][j] && s1[i-1] == s3[k])
		}
	}
	return dp[len1][len2]
}
// @lc code=end

