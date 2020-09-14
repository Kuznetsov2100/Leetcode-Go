/*
 * @lc app=leetcode.cn id=1572 lang=golang
 *
 * [1572] 矩阵对角线元素的和
 */

// @lc code=start
func diagonalSum(mat [][]int) int {
	var res int
	m := len(mat)
	for i , j:= 0, m-1; i < m && j >= 0; i, j = i+1, j-1 {
		res += mat[i][i]+mat[i][j]
	}
	if m % 2 == 1 {
		return res-mat[(m-1)/2][(m-1)/2]
	}
	return res
}
// @lc code=end

