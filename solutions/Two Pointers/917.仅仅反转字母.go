/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
func reverseOnlyLetters(S string) string {
	chars := []byte(S)
	i, j := 0, len(chars)-1
	for i < j {
		if !isAlpha(chars[i]) {
			i++
		}
		if !isAlpha(chars[j]) {
			j--
		}
		if isAlpha(chars[i]) && isAlpha(chars[j]) {
			chars[i], chars[j] = chars[j], chars[i]
			i++
			j--
		}
	}
	return string(chars)
}

func isAlpha (x byte) bool {
	if 'a' <= x && x <= 'z' || 'A' <= x && x <= 'Z' {
		return true
	}
	return false
}
// @lc code=end

