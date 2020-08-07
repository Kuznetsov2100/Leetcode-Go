/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	isValid := func(r, c int, num byte) bool {
		for i := 0; i < 9; i++ {
			if board[r][i] == num && i != c {
				return false
			}
			if board[i][c] == num && i != r {
				return false
			}
			if row, col := ((r/3)*3+i/3),((c/3)*3+i%3); board[row][col] == num && row != r && col != c  {
				return false
			}
		}
		return  true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if !isValid(i,j,board[i][j]) {
					return false
				}
			}
		}
	}
	return true	
}

// @lc code=end

