/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */

// @lc code=start
func rangeBitwiseAnd(m int, n int) int {
	// m,n每次都右移一位，知道m等于n,shift记录右移次数，
	// 最后的m即为原[m,n]之间的公共前缀，右边补shift个0即为结果
	shift := 0
	for m < n {
		m = m >> 1
		n = n >> 1
		shift++
	}
	return m << shift
}
// @lc code=end

