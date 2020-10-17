/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 */

// @lc code=start
func largestPerimeter(A []int) int {
	// 问题本质： 找到数组中能组成三角形的最大的三个数
	sort.Ints(A)
	for i := len(A)-3; i >= 0; i-- {
		if A[i]+ A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}
// @lc code=end

