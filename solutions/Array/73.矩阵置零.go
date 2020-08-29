/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int)  {
	row_zero := make(map[int]bool)
	col_zero := make(map[int]bool)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				row_zero[i] = true
				col_zero[j] = true
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if row_zero[i] || col_zero[j] {
				matrix[i][j] = 0
			}
		}
	}
}
// @lc code=end

