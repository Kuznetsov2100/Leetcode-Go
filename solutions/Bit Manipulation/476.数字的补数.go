/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	// 将cnt设为与num二进制位数相同且是该位数下最大的二进制数
	// for example: num 101, cnt 111
	cnt := 1
	for cnt < num { 
		cnt <<= 1
		cnt += 1
	}
	return num ^ cnt
}
// @lc code=end

