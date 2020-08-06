/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
import "strings"
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for !isAlphaNum(s[i]) && i < j {
			i++
		}
		for !isAlphaNum(s[j]) && i < j {
			j--
		}
		if i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}

func isAlphaNum(x byte) bool {
	if x >= '0' && x <= '9' || x >= 'a' && x <= 'z' {
		return true
	} 
	return false
}
// @lc code=end

