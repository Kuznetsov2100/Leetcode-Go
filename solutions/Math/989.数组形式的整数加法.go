/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 */

// @lc code=start
func addToArrayForm(A []int, K int) []int {
	res := A
	for K > 0 {
		res = plusOne(res)
		K--
	}
	return res
}
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

