/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	// dp[i][j]的含义：字串[i..j]能否形成回文 ,true为可以，false为不可以
	dp := make([][]bool,n)
	for i := range dp {
		dp[i] = make([]bool,n)
	}
	res := s[0:1] // res初始化为s的第一个字符
	for i := n-1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j { // 长度为1的回文子串
				dp[i][j] = true
			} else if j - i == 1 && s[i] == s[j] { // 长度为2的回文子串
				dp[i][j] = true
			} else if s[i] == s[j] && dp[i+1][j-1] { // 字串[i+1..j-1]形成了回文，当s[i] == s[j]时，字串[i..j]形成了回文
				dp[i][j] = true
			}
			if dp[i][j] && j-i+1 > len(res) { // 如果回文[i..j]长度超过res，更新res
				res = s[i:j+1]
			}
		}
	}
	return res
}
// @lc code=end

