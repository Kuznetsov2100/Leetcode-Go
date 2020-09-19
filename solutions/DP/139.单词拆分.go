/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// dp[i]表示字符串s的前i个字符可以被空格拆分为一个或多个在字典中出现的单词
	// [0, i-1]区间子串能否被成功拆分，取决于两部分
	// [0, j-1]区间能被划分，[j,i-1]部分存在于wordDict中， 则dp[i] = true
	// 指针j用于切割字符串
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true // base case
	seen := make(map[string]bool) // wordDict hash set
	for _, word := range wordDict {
		seen[word] = true
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && seen[s[j:i]] {
				dp[i] = true
				break // 一旦确认了子串[0,i-1]能被成功拆分，则无需再用指针j切割[0,i-1]部分
			}
		}
	}
	return dp[n]
}
// @lc code=end

