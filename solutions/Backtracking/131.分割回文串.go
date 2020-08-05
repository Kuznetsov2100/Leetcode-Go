/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	var res [][]string
	if s == "" {
		return res
	}
	isPalindrome := func(s string) bool {
		i, j := 0, len(s)-1
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	var backtrack func(str string, track []string)
	backtrack = func(str string, track []string) {
		if str == "" && len(track) > 0 {
			tmp := make([]string, len(track))
			copy(tmp,track)
			res = append(res, tmp)
			return
		}
		for i := 1; i <= len(s); i++ {
			if i > len(str) {
				return
			}
			if i > 1 && !isPalindrome(str[:i]) {
				continue
			}
			track = append(track, str[:i])
			backtrack(str[i:],track)
			track = track[:len(track)-1]
		}
	}
	backtrack(s,[]string{})
	return res
}
// @lc code=end

