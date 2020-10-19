/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	res, i, last, n := 0, 0, 0, len(s)
	for i < n {
		c := s[i]
		cnt := 0
		for i < n && s[i] == c {
			i++
			cnt++
		}
		res += min(cnt, last)
		last = cnt
	}
	return res	
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

