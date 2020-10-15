/*
 * @lc app=leetcode.cn id=944 lang=golang
 *
 * [944] 删列造序
 */

// @lc code=start
func minDeletionSize(A []string) int {
	var res int
	lengths := len(A[0])
	lengthA := len(A)
	for i := 0; i < lengths; i++ {
		for j := 0; j < lengthA-1; j++ {
			if A[j][i] > A[j+1][i] {
				res++
				break
			}
		}
	}
	return res
}
// @lc code=end

