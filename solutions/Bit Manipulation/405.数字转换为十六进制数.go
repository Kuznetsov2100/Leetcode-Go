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
	// 如果num是负数，右移不能保证num最终为0，因为num是32位int,对应8位十六进制
	// 所以使用 len(res) < 8 来保证循环终止
	for num != 0 && len(res) < 8 { 
		res = string(hex[num & 0xf]) + res  // num & 0xf 获取num的低4位，根据hex得到十六进制表示
		num >>= 4 // 如果num是正整数，右移左边补0，如果num是负数，右移左边补1
	}
	return res
}
// @lc code=end

