/*
 * @lc app=leetcode.cn id=1009 lang=golang
 *
 * [1009] 十进制整数的反码
 */

// @lc code=start
func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	cnt := 1
	for cnt < N { 
		cnt <<= 1
		cnt += 1
	}
	return N ^ cnt
}
// @lc code=end

