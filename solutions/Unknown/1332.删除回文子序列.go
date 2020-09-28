/*
 * @lc app=leetcode.cn id=1332 lang=golang
 *
 * [1332] 删除回文子序列
 */

// @lc code=start
func removePalindromeSub(s string) int {
	if s == "" {
		return 0
	}
	if isPalindrome(s) {
		return 1
	}
	return 2
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
// @lc code=end

