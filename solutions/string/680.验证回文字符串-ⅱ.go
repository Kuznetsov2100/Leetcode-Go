/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

// @lc code=start
func validPalindrome(s string) bool {
	if ok, indexi, indexj := isPalindrome(s); ok {
		return true
	} else {
		if ok1, _, _ := isPalindrome(s[:indexi]+s[indexi+1:]); ok1 {
			return true
		} else {
			if ok2, _, _ := isPalindrome(s[:indexj]+s[indexj+1:]); ok2 {
				return true
			}
			return false
		}
	}
}

func isPalindrome(s string) (bool, int, int) {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false,i,j
		}
		i++
		j--
	}
	return true,-1,-1
}
// @lc code=end

