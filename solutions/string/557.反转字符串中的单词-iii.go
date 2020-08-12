/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
import "strings"
func reverseWords(s string) string {
	str := strings.Split(s, " ")
	for i := range str {
		str[i] = reverseString(str[i])
	}
	return strings.Join(str, " ")
}

func reverseString(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i,j = i+1,j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
// @lc code=end

