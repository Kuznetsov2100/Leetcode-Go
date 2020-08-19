/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] 数字转换为十六进制数
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdef"
	var res string
	for num != 0 && len(res) < 8 {
		res = string(hex[num & 0xf]) + res
		num >>= 4
	}
	return res
}
// @lc code=end

