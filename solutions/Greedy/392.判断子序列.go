/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	indexS, indexT := 0, 0
	for indexS < m && indexT < n {
		if s[indexS] == t[indexT] {
			indexS++
		}
		indexT++
	}
	return indexS == m
}
// @lc code=end

