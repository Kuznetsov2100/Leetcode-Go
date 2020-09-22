/*
 * @lc app=leetcode.cn id=1342 lang=golang
 *
 * [1342] 将数字变成 0 的操作次数
 */

// @lc code=start
func numberOfSteps (num int) int {
	var res int
	for num > 0 {
		res++
		if num % 2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
	}
	return res
}
// @lc code=end

