/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
import "strconv"
func convertToBase7(num int) string {
	var res string
	sign := 1
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = -num
		sign = -1
	}
	for num != 0 {
		res = strconv.Itoa(num % 7) + res
		num /= 7
	}
	if sign == -1 {
		return "-" + res
	}
	return res
}
// @lc code=end

