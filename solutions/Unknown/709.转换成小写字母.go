/*
 * @lc app=leetcode.cn id=709 lang=golang
 *
 * [709] 转换成小写字母
 */

// @lc code=start
func toLowerCase(str string) string {
	s := []byte(str)
	for i := range s {
		s[i] |= ' '
	}
	return string(s)
}
// @lc code=end

