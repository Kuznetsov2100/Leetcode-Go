/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
import "strings"
func reverseWords(s string) string {
	str := strings.Fields(s)
	for i,j := 0, len(str)-1; i < j; i,j = i+1,j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return strings.Join(str, " ")
}
// @lc code=end

