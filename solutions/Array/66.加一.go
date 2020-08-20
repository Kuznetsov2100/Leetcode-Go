/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return []int{1}
	}
	digits[n-1]++	
	if digits[n-1] == 10 {
		return append(plusOne(digits[:n-1]),0)
	} 
	return digits
}
// @lc code=end