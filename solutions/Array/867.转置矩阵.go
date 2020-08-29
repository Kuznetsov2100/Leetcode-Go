/*
 * @lc app=leetcode.cn id=867 lang=golang
 *
 * [867] 转置矩阵
 */

// @lc code=start
func transpose(A [][]int) [][]int {
	row, col := len(A), len(A[0])
	res := make([][]int, col)
	for i := range res {
		res[i] = make([]int, row)
	}
	for i := 0; i < row*col; i++ {
		res[i%col][i/col] = A[i/col][i%col]
	}
	return res
}
// @lc code=end

