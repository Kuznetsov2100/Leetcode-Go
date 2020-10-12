/*
 * @lc app=leetcode.cn id=1304 lang=golang
 *
 * [1304] 和为零的N个唯一整数
 */

// @lc code=start
func sumZero(n int) []int {
	res := make([]int, 0, n)
	sum := 0
	for i := 1; i < n; i++ {
		res = append(res, i)
		sum += i
	}
	res = append(res, -sum)
	return res
}
// @lc code=end

