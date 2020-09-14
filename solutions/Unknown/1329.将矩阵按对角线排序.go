/*
 * @lc app=leetcode.cn id=1329 lang=golang
 *
 * [1329] 将矩阵按对角线排序
 */

// @lc code=start
func diagonalSort(mat [][]int) [][]int {
	// 对角线冒泡排序
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; i+k < m && j+k < n; k++ {
				if mat[i][j] > mat[i+k][j+k] {
					mat[i][j], mat[i+k][j+k] = mat[i+k][j+k], mat[i][j]
				}
			}
		}
	}
	return mat
}
// @lc code=end

