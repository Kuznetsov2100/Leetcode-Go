/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	if num >= 0 && num <= 9 {
		return num
	}
	if x := num % 10; x != 0 {
		return addDigits(x + addDigits(num-x))
	} else {
		return addDigits(num/10)
	}
}
// @lc code=end

