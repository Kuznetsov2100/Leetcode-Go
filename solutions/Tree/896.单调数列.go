/*
 * @lc app=leetcode.cn id=896 lang=golang
 *
 * [896] 单调数列
 */

// @lc code=start
func isMonotonic(A []int) bool {
	positive,negative := false,false
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			positive = true
		} else if A[i] > A[i+1] {
			negative = true
		}
		if positive && negative {
			return false
		}
	}
	return true
}
// @lc code=end

