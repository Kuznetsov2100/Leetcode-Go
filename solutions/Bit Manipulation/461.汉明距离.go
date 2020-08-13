/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	n := x ^ y
	var res int
	for n != 0 {
		n = n & (n-1)
		res++
	}
	return res
}
// @lc code=end

