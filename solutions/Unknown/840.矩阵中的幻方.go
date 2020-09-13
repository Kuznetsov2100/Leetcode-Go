/*
 * @lc app=leetcode.cn id=840 lang=golang
 *
 * [840] 矩阵中的幻方
 */

// @lc code=start
func numMagicSquaresInside(grid [][]int) int {
	var res int
	m, n := len(grid), len(grid[0])
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if isValidMagicSquare(i, j, grid) {
				res++
			}
		}
	}
	return res
}

func isValidMagicSquare(r, c int, grid [][]int) bool {
	var rowsum, colsum, diagonalsum int
	seen := make([]bool, 10)
	for i := 0; i < 9; i++ {
		value := grid[r-1+i/3][c-1+i%3]
		if value > 9 || value < 1 || seen[value] {
			return false
		}
		seen[value] = true
		if i == 0 || i == 1 || i == 2 {
			rowsum += value
		}
		if i == 0 || i == 3 || i == 6 {
			colsum += value
		}
		if i == 0 || i == 4 || i == 8 {
			diagonalsum += value
		}		
	}
	return  rowsum == colsum && rowsum == diagonalsum
}
// @lc code=end

