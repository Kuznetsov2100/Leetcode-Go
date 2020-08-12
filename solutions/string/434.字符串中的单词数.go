/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func countSegments(s string) int {
	var res int
	for i := 0; i < len(s); {
		oldi := i
		for ;i < len(s) && s[i] != ' '; i++ {

		}
		if i > oldi {
			res++
		}
		for ; i < len(s) && s[i] == ' '; i++ {

		}
	}
	return res
}
// @lc code=end

