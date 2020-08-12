/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */

// @lc code=start
import "strings"
func detectCapitalUse(word string) bool {
	if word == strings.ToUpper(word) || word == word[0:1]+strings.ToLower(word[1:]) {
		return true
	}
	return false
}
// @lc code=end

