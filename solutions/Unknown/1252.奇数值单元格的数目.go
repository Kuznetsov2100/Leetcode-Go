/*
 * @lc app=leetcode.cn id=1252 lang=golang
 *
 * [1252] 奇数值单元格的数目
 */

// @lc code=start
func oddCells(n int, m int, indices [][]int) int {
	var res int
	rowsCount := make([]int, n)
	colsCount := make([]int, m)
	for _, indice := range indices {
		rowsCount[indice[0]]++
		colsCount[indice[1]]++
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (rowsCount[i]+colsCount[j]) % 2 == 1 {
				res++
			}
		}
	}
	return res
}
// @lc code=end

