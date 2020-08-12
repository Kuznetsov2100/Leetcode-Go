/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
import "strings"
func lengthOfLastWord(s string) int {
	str := strings.Split(s, " ")
	var nonempty []string
	for i := range str {
		if len(str[i]) > 0 {
			nonempty = append(nonempty, str[i])
		}
	}
	if len(nonempty) == 0 {
		return 0
	}
	return len(nonempty[len(nonempty)-1])
}
// @lc code=end

