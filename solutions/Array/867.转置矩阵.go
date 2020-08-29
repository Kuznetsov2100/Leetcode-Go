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
		r , c := i/col, i%col
		res[c][r] = A[r][c]
	}
	return res
}
// @lc code=end

