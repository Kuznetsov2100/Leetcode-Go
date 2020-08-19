/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
import "strings"
func convertToTitle(n int) string {
	var res string
	for n != 0 {
		n--
		res = string(n % 26 + 'A') + res
		n = n/26
	}
	return res
}
// @lc code=end

